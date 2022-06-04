import React from "react";
import "./App.css";
import { GoogleMap, LoadScript, Marker } from "@react-google-maps/api";
import mapStyles from "./mapStyles.json";

const containerStyle = {
  width: "min(1600px, 90vw)",
  height: "min(800px, 90vh)",
};

const center = {
  lat: 40.710572745064215,
  lng: -73.95191000508696,
};

function pointToMarker(p, i) {
  return (
    <Marker
      key={i}
      title={`(${p.lat}, ${p.lng})`}
      position={{ lat: p.lat, lng: p.lng }}
      opacity={0.5}
    />
  );
}

function MyMap() {
  const apiKey = process.env.REACT_APP_GOOGLE_MAPS_API_KEY;
  const [points, setPoints] = React.useState([]);
  const [clickedPoints, setClickedPoints] = React.useState([]);
  React.useEffect(() => {
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
      setPoints(
        body.driverLocations.map((dl) => ({
          lat: dl.currentLocation.latitude,
          lng: dl.currentLocation.longitude,
        }))
      );
      console.log("body", body);
    };

    getDriverLocations().catch(console.error);
  });
  return (
    <LoadScript googleMapsApiKey={apiKey}>
      <GoogleMap
        onClick={(e) => {
          const lat = e.latLng.lat();
          const lng = e.latLng.lng();
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
  );
}

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <div style={{ marginTop: "20px", marginBottom: "100px" }}>
          <MyMap />
        </div>
      </header>
    </div>
  );
}

export default App;
