import React, { ReactElement } from "react";
import "./App.css";
import { GoogleMap, LoadScript } from "@react-google-maps/api";
import Marker from "./Marker";
import mapStyles from "./mapStyles.json";
import { ThemeProvider, createTheme } from "@mui/material/styles";
import {
  Box,
  FormControlLabel,
  FormGroup,
  Grid,
  Switch,
  Typography,
} from "@mui/material";
import { styled } from "@mui/material/styles";
import DriverList from "./PointList";
import Paper from "@mui/material/Paper";
import { useEffectOnce } from "usehooks-ts";
import {
  addDriverLocationToState,
  DriverLocation,
  driverLocationsToState,
  getDriverLocationsFromState,
  LatLng,
  NormalizedDriverLocations,
  SearchResult,
} from "./types";
import { faker } from "@faker-js/faker";
import listDrivers from "./request/listDrivers";
import getNearestDrivers from "./request/getNearestDrivers";
import updateDriverLocations from "./request/updateDriverLocations";

const darkTheme = createTheme({
  palette: {
    mode: "dark",
  },
});

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
  const [focusedDriverId, setFocusedDriverId] = React.useState<string>("");
  const [queryMode, setQueryMode] = React.useState<boolean>(false);
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
  const [searchResults, setSearchResults] = React.useState<SearchResult[]>([]);

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

  const getNearbyDrivers = async (e: google.maps.MapMouseEvent) => {
    const lat = e.latLng?.lat() ?? 0;
    const lng = e.latLng?.lng() ?? 0;

    const res = await getNearestDrivers({
      latitude: lat,
      longitude: lng,
    } as LatLng);
    setSearchResults(res?.results ?? []);
  };

  const addNewDriver = (e: google.maps.MapMouseEvent) => {
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
  };

  return (
    <Grid container spacing={2}>
      <Grid item xs={3}>
        <FormGroup>
          <FormControlLabel
            control={
              <Switch
                checked={queryMode}
                onChange={() => setQueryMode(!queryMode)}
              />
            }
            label="Mode"
          />
        </FormGroup>
      </Grid>
      <Grid item xs={9}>
        {queryMode ? (
          <Typography>Click anywhere to rank nearby drivers</Typography>
        ) : (
          <Typography>Click to add new drivers</Typography>
        )}
      </Grid>
      <Grid item xs={3}>
        <Item>
          <DriverList
            buildHandleMouseOver={(driverId: string) => () =>
              setFocusedDriverId(driverId)}
            buildHandleMouseOut={(driverId: string) => () =>
              setFocusedDriverId("")}
            queryMode={queryMode}
            onUpload={() =>
              updateDriverLocations(
                getDriverLocationsFromState(newDriverLocationsState)
              )
            }
            driverLocations={newDriverLocations}
          />
        </Item>
      </Grid>
      <Grid item xs={9}>
        <Item>
          <LoadScript googleMapsApiKey={apiKey}>
            <GoogleMap
              onClick={queryMode ? getNearbyDrivers : addNewDriver}
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
                  isNear={searchResults
                    .map((sr) => sr.driver.driverId)
                    .includes(dl.driverId)}
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
    <ThemeProvider theme={darkTheme}>
      <Item
        sx={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          height: "100vh",
          px: {
            xs: "5px",
            sm: "10px",
            md: "20px",
          },
        }}
      >
        <MyMap />
      </Item>
    </ThemeProvider>
  );
}

export default App;
