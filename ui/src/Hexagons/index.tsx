import React from "react";
import { getPolygons, getPolygonsOutput } from "../getPolygons";
import { LatLng } from "../types";
import { Polygon } from "@react-google-maps/api";

interface HexagonsProps {
  polyOut: getPolygonsOutput;
}

export function buildOptions(k: number): google.maps.PolygonOptions {
  const fillColor = k === 0 ? "lightblue" : k === 1 ? "lightblue" : "lightblue";
  return {
    fillColor,
    fillOpacity: 0.3,
    strokeColor: "red",
    strokeOpacity: 1,
    strokeWeight: 1,
    clickable: false,
    draggable: false,
    editable: false,
    geodesic: false,
    zIndex: 1,
  } as google.maps.PolygonOptions;
}

export default function Hexagons({ polyOut }: HexagonsProps) {
  const onLoad = (polygon: google.maps.Polygon) => {
    console.log("polygon: ", polygon);
  };
  return (
    <>
      {polyOut?.ring0.map((points: LatLng[]) => (
        <Polygon
          onLoad={onLoad}
          paths={pointsToPaths(points)}
          options={buildOptions(0)}
        />
      )) ?? null}
      {polyOut?.ring1.map((points: LatLng[]) => (
        <Polygon
          onLoad={onLoad}
          paths={pointsToPaths(points)}
          options={buildOptions(1)}
        />
      )) ?? null}
      {polyOut?.ring2.map((points: LatLng[]) => (
        <Polygon
          onLoad={onLoad}
          paths={pointsToPaths(points)}
          options={buildOptions(2)}
        />
      )) ?? null}
    </>
  );
}

export function pointsToPaths(points: LatLng[]): google.maps.LatLngLiteral[] {
  return points.map(
    (p: LatLng) =>
      ({ lat: p.latitude, lng: p.longitude } as google.maps.LatLngLiteral)
  );
}
