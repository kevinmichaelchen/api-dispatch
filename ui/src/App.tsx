import React from "react";
import "./App.css";
import { GoogleMap, LoadScript, Marker } from "@react-google-maps/api";
import mapStyles from "./mapStyles.json";
import { Grid } from "@mui/material";
import { styled } from "@mui/material/styles";
import Box from "@mui/material/Box";
import Paper from "@mui/material/Paper";
import { useEffectOnce } from "usehooks-ts";

const Item = styled(Paper)(({ theme }) => ({
  backgroundColor: theme.palette.mode === "dark" ? "#1A2027" : "#fff",
  ...theme.typography.body2,
  padding: theme.spacing(1),
  textAlign: "center",
  color: theme.palette.text.secondary,
}));

const containerStyle = {
  width: "100%",
  height: "min(800px, 100vh)",
};

const center = {
  lat: 40.710572745064215,
  lng: -73.95191000508696,
};

interface Point {
  lat: number;
  lng: number;
}

function pointToMarker(p: Point, i: number) {
  return (
    <Marker
      key={i}
      title={`(${p.lat}, ${p.lng})`}
      position={{ lat: p.lat, lng: p.lng }}
      opacity={0.5}
    />
  );
}

interface LatLng {
  latitude: number;
  longitude: number;
}

interface DriverLocation {
  currentLocation: LatLng;
  driverId: string;
  id: string;
  mostRecentHeartbeat: string;
}

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
        <Item>Cached points go here</Item>
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
              {points.map((p, i) => pointToMarker(p, i))}
              {clickedPoints.map((p, i) => pointToMarker(p, i))}
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
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "100vh",
        backgroundColor: "#282c34",
        paddingTop: "10px",
      }}
    >
      <MyMap />
    </div>
  );
}

export default App;
