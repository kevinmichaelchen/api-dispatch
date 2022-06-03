import logo from "./logo.svg";
import "./App.css";
import { GoogleMap, LoadScript } from "@react-google-maps/api";
import mapStyles from "./mapStyles.json";

const containerStyle = {
  width: "min(1200px, 80vw)",
  height: "min(600px, 80vh)",
};

const center = {
  lat: 40.730572745064215,
  lng: -73.95191000508696,
};

function MyMap() {
  const apiKey = process.env.REACT_APP_GOOGLE_MAPS_API_KEY;
  console.log("apiKey", apiKey);
  return (
    <LoadScript googleMapsApiKey={apiKey}>
      <GoogleMap
        options={{
          styles: mapStyles,
        }}
        mapContainerStyle={containerStyle}
        center={center}
        zoom={10}
      >
        {/* Child components, such as markers, info windows, etc. */}
        <></>
      </GoogleMap>
    </LoadScript>
  );
}

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <div style={{ marginTop: "100px", marginBottom: "200px" }}>
          <MyMap />
        </div>
      </header>
    </div>
  );
}

export default App;
