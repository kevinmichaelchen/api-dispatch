import { LatLng } from "./types";
import { geoToH3, H3Index, h3ToGeoBoundary, kRing } from "h3-js";

export interface getPolygonsOutput {
  ring0: LatLng[][];
  ring1: LatLng[][];
  ring2: LatLng[][];
}

export function getPolygons(l: LatLng, res: number): getPolygonsOutput {
  const index = geoToH3(l.latitude, l.longitude, res);
  console.log("index", index);

  const indexes0 = kRing(index, 0);
  const indexes1 = kRing(index, 1);
  const indexes2 = kRing(index, 2);

  return {
    ring0: indexes0.map(getIndexVertexes),
    ring1: indexes1.map(getIndexVertexes),
    ring2: indexes2.map(getIndexVertexes),
  } as getPolygonsOutput;
}

function getIndexVertexes(index: H3Index): LatLng[] {
  const arr = h3ToGeoBoundary(index);
  return arr.map(
    (p: number[]) =>
      ({
        latitude: p[0],
        longitude: p[1],
      } as LatLng)
  );
}
