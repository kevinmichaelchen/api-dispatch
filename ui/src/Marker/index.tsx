import React from "react";
import { renderToString } from "react-dom/server";
import { Point } from "../types";
import { Marker } from "@react-google-maps/api";
import { DirectionsCarFilled } from "@mui/icons-material";

function getIconPath(element: React.ReactElement): string {
  const iconString = renderToString(element);
  const parser = new DOMParser();
  const svgDoc = parser.parseFromString(iconString, "image/svg+xml");
  return svgDoc.querySelector("path")?.getAttribute("d") as string;
}

interface MyMarkerProps {
  point: Point;
  handleMouseOver?: (e: google.maps.MapMouseEvent) => void;
  cached?: boolean;
}

export default function MyMarker(props: MyMarkerProps) {
  const { cached, point: p, handleMouseOver } = props;
  const color = cached ? "orange" : "#FFD700";
  return (
    <Marker
      title={`(${p.lat}, ${p.lng})`}
      position={{ lat: p.lat, lng: p.lng }}
      opacity={1}
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
