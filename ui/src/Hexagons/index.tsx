import { getPolygons } from "../getPolygons";
import { LatLng } from "../types";
import { Polygon } from "@react-google-maps/api";

interface HexagonsProps {
  pickupLocation: LatLng;
  resolution: number;
}

function buildOptions(k: number): google.maps.PolygonOptions {
  const fillColor = k === 0 ? "lightblue" : k === 1 ? "lightblue" : "lightblue";
  return {
    fillColor,
    fillOpacity: 1,
    strokeColor: "red",
    strokeOpacity: 1,
    strokeWeight: 2,
    clickable: false,
    draggable: false,
    editable: false,
    geodesic: false,
    zIndex: 1,
  };
}

export default function Hexagons(props: HexagonsProps) {
  const { pickupLocation, resolution } = props;
  const out = getPolygons(pickupLocation, resolution);
  return (
    <>
      {out.ring0.map((points: LatLng[]) => (
        <Polygon paths={pointsToPaths(points)} options={buildOptions(0)} />
      ))}
      {out.ring1.map((points: LatLng[]) => (
        <Polygon paths={pointsToPaths(points)} options={buildOptions(1)} />
      ))}
      {out.ring2.map((points: LatLng[]) => (
        <Polygon paths={pointsToPaths(points)} options={buildOptions(2)} />
      ))}
    </>
  );
}

function pointsToPaths(points: LatLng[]): google.maps.LatLngLiteral[] {
  return points.map(
    (p) => ({ lat: p.latitude, lng: p.longitude } as google.maps.LatLngLiteral)
  );
}
