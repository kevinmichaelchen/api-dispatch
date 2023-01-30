import React from "react";
import { renderToString } from "react-dom/server";
import { Marker } from "@react-google-maps/api";
import { DirectionsCarFilled } from "@mui/icons-material";
import { DriverLocation, LatLng } from "../types";

function getIconPath(element: React.ReactElement): string {
  const iconString = renderToString(element);
  const parser = new DOMParser();
  const svgDoc = parser.parseFromString(iconString, "image/svg+xml");
  return svgDoc.querySelector("path")?.getAttribute("d") as string;
}

interface MyMarkerProps {
  /**
   * The marker's location
   */
  location: LatLng;
  handleMouseOver?: (e: google.maps.MapMouseEvent) => void;
  handleMouseOut?: (e: google.maps.MapMouseEvent) => void;
  handleClick?: (e: google.maps.MapMouseEvent) => void;

  /**
   * Whether the marker is for the requested pickup location
   */
  pickupLocation?: boolean;
  /**
   * Whether the marker is for a new driver location that will be created in bulk.
   */
  cached?: boolean;
  /**
   * Whether the marker represents a nearby driver to the requested pickup location.
   */
  isNear?: boolean;
}

export default function MyMarker(props: MyMarkerProps) {
  const {
    pickupLocation,
    cached,
    isNear,
    location: p,
    handleClick,
    handleMouseOver,
    handleMouseOut,
  } = props;
  let color = "#FFD700";
  if (isNear) color = "green";
  if (cached) color = "orange";
  if (pickupLocation) color = "purple";
  if (pickupLocation)
    console.log("rendering marker for pickup", pickupLocation, p);
  return (
    <Marker
      onUnmount={(marker: google.maps.Marker) => console.log("marker unmount")}
      title={`(${p.latitude}, ${p.longitude})`}
      position={{ lat: p.latitude, lng: p.longitude }}
      opacity={1}
      onClick={handleClick}
      onMouseOver={handleMouseOver}
      onMouseOut={handleMouseOut}
      icon={{
        path: getIconPath(<DirectionsCarFilled />),
        fillColor: color,
        fillOpacity: 0.9,
        scale: 0.7,
        strokeColor: color,
        strokeWeight: 0.5,
      }}
    />
  );
}
