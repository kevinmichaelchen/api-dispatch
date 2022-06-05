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
  driverLocation: DriverLocation;
  handleMouseOver?: (e: google.maps.MapMouseEvent) => void;
  handleClick?: (e: google.maps.MapMouseEvent) => void;
  cached?: boolean;
}

export default function MyMarker(props: MyMarkerProps) {
  const { cached, driverLocation, handleClick, handleMouseOver } = props;
  const color = cached ? "orange" : "#FFD700";
  const p = driverLocation.currentLocation;
  return (
    <Marker
      title={`(${p.latitude}, ${p.longitude})`}
      position={{ lat: p.latitude, lng: p.longitude }}
      opacity={1}
      onClick={handleClick}
      onMouseOver={handleMouseOver}
      icon={{
        path: getIconPath(<DirectionsCarFilled />),
        fillColor: color,
        fillOpacity: 0.9,
        scale: 0.7,
        strokeColor: color,
        strokeWeight: 1,
      }}
    />
  );
}
