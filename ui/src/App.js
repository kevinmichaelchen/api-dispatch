import React from "react";
import logo from "./logo.svg";
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

function MyMap() {
  const apiKey = process.env.REACT_APP_GOOGLE_MAPS_API_KEY;
  const [points, setPoints] = React.useState([]);
  return (
    <LoadScript googleMapsApiKey={apiKey}>
      <GoogleMap
        onClick={(e) => {
          const lat = e.latLng.lat();
          const lng = e.latLng.lng();
          console.log(`clicked (${lat}, ${lng})`);
          setPoints([...points, { lat, lng }]);
        }}
        options={{
          styles: mapStyles,
        }}
        mapContainerStyle={containerStyle}
        center={center}
        zoom={13}
      >
        {points.map((p, i) => (
          <Marker key={i} position={{ lat: p.lat, lng: p.lng }} />
        ))}
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
