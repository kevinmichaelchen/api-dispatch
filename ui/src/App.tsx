import React, { ReactElement } from "react";
import "./App.css";
import { GoogleMap, LoadScript } from "@react-google-maps/api";
import Marker from "./Marker";
import mapStyles from "./mapStyles.json";
import { Box, Grid } from "@mui/material";
import { styled } from "@mui/material/styles";
import PointList from "./PointList";
import Paper from "@mui/material/Paper";
import { useEffectOnce } from "usehooks-ts";
import { DriverLocation, Point } from "./types";

const Item = styled(Paper)(({ theme }) => ({
  backgroundColor: theme.palette.mode === "dark" ? "#1A2027" : "#fff",
  ...theme.typography.body2,
  padding: theme.spacing(1),
  textAlign: "center",
  color: theme.palette.text.secondary,
}));

const containerStyle = {
  width: "100%",
  height: "min(800px, 90vh)",
};

const center = {
  lat: 40.710572745064215,
  lng: -73.95191000508696,
};

function MyMap() {
  const apiKey = process.env.REACT_APP_GOOGLE_MAPS_API_KEY as string;
  const [points, setPoints] = React.useState<Point[]>([]);
  const [clickedPoints, setClickedPoints] = React.useState<Point[]>([]);

  useEffectOnce(() => {
    const getDriverLocations = async () => {
      const response = await fetch(
        "http://localhost:8081/coop.drivers.dispatch.v1beta1.DispatchService/ListDrivers",
        {
          method: "POST",
          mode: "cors",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ page_size: 1000 }),
        }
      );
      console.log("response", response);
      const body = await response.json();
      console.log("body", body);
      setPoints(
        body.driverLocations.map((dl: DriverLocation) => ({
          lat: dl.currentLocation.latitude,
          lng: dl.currentLocation.longitude,
        }))
      );
    };

    getDriverLocations().catch(console.error);
  });

  return (
    <Grid container spacing={2}>
      <Grid item xs={3}>
        <Item>
          <PointList points={clickedPoints} />
        </Item>
      </Grid>
      <Grid item xs={9}>
        <Item>
          <LoadScript googleMapsApiKey={apiKey}>
            <GoogleMap
              onClick={(e: google.maps.MapMouseEvent) => {
                const lat = e.latLng?.lat() ?? 0;
                const lng = e.latLng?.lng() ?? 0;
                console.log(`clicked (${lat}, ${lng})`);
                setClickedPoints([...clickedPoints, { lat, lng }]);
              }}
              options={{
                styles: mapStyles,
              }}
              mapContainerStyle={containerStyle}
              center={center}
              zoom={13}
            >
              {points.map((p: Point, i: number) => (
                <Marker key={i} point={p} />
              ))}
              {clickedPoints.map((p: Point, i: number) => (
                <Marker key={i} point={p} cached />
              ))}
              <></>
            </GoogleMap>
          </LoadScript>
        </Item>
      </Grid>
    </Grid>
  );
}

function App() {
  return (
    <Box
      sx={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "100vh",
        backgroundColor: "#282c34",
        px: {
          xs: "5px",
          sm: "10px",
          md: "20px",
        },
      }}
    >
      <MyMap />
    </Box>
  );
}

export default App;
