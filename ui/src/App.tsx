import React, { ReactElement } from "react";
import "./App.css";
import { GoogleMap, LoadScript } from "@react-google-maps/api";
import Marker from "./Marker";
import mapStyles from "./mapStyles.json";
import { Box, Grid } from "@mui/material";
import { styled } from "@mui/material/styles";
import NewDriverList from "./PointList";
import Paper from "@mui/material/Paper";
import { useEffectOnce } from "usehooks-ts";
import {
  addDriverLocationToState,
  DriverLocation,
  driverLocationsToState,
  getDriverLocationsFromState,
  LatLng,
  NormalizedDriverLocations,
} from "./types";
import { faker } from "@faker-js/faker";
import listDrivers from "./request/listDrivers";
import getNearestDrivers from "./request/getNearestDrivers";

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

function newDriverName(): string {
  const firstName = faker.name.firstName();
  const middleName = faker.name.middleName();
  const lastName = faker.name.lastName();
  const fullName = `${firstName}-${middleName}-${lastName}`;
  return fullName.replaceAll(" ", "-");
}

function MyMap() {
  const apiKey = process.env.REACT_APP_GOOGLE_MAPS_API_KEY as string;
  const [driverLocationsState, setDriverLocationsState] =
    React.useState<NormalizedDriverLocations>({
      byId: {},
      allIds: [],
    } as NormalizedDriverLocations);
  const [newDriverLocationsState, setNewDriverLocationsState] =
    React.useState<NormalizedDriverLocations>({
      byId: {},
      allIds: [],
    } as NormalizedDriverLocations);

  const newDriverLocations = getDriverLocationsFromState(
    newDriverLocationsState
  );
  const driverLocations = getDriverLocationsFromState(driverLocationsState);

  useEffectOnce(() => {
    const getDriverLocations = async () => {
      const driverLocations = await listDrivers();
      setDriverLocationsState(driverLocationsToState(driverLocations));
    };

    getDriverLocations().catch(console.error);
  });

  return (
    <Grid container spacing={2}>
      <Grid item xs={3}>
        <Item>
          <NewDriverList driverLocations={newDriverLocations} />
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
                const dl = {
                  driverId: newDriverName(),
                  currentLocation: {
                    latitude: lat,
                    longitude: lng,
                  },
                } as DriverLocation;
                setNewDriverLocationsState(
                  addDriverLocationToState(newDriverLocationsState, dl)
                );
              }}
              options={{
                styles: mapStyles,
              }}
              mapContainerStyle={containerStyle}
              center={center}
              zoom={13}
            >
              {driverLocations.map((dl: DriverLocation, i: number) => (
                <Marker
                  key={i}
                  driverLocation={dl}
                  // TODO move this to Map.onClick, which will AddNewDriver if in Mutation/Ingest Mode, or query/dispatch if in Query/Dispatch Mode
                  handleClick={(e) =>
                    getNearestDrivers({
                      latitude: e.latLng?.lat() ?? 0,
                      longitude: e.latLng?.lng() ?? 0,
                    })
                  }
                />
              ))}
              {newDriverLocations.map((dl: DriverLocation, i: number) => (
                <Marker key={i} driverLocation={dl} cached />
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
