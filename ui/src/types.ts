export interface LatLng {
  latitude: number;
  longitude: number;
}

export interface DriverLocation {
  currentLocation: LatLng;
  driverId: string;
  id: string;
  mostRecentHeartbeat: string;
}

export interface NormalizedDriverLocations {
  byId: { [key: string]: DriverLocation };
  allIds: string[];
}

export function driverLocationsToState(
  driverLocations: DriverLocation[]
): NormalizedDriverLocations {
  return {
    allIds: driverLocations.map((dl) => dl.driverId),
    byId: driverLocations.reduce((acc: any, dl: DriverLocation) => {
      acc[dl.driverId] = dl;
      return acc;
    }, {}),
  };
}

export function addDriverLocationToState(
  s: NormalizedDriverLocations,
  dl: DriverLocation
): NormalizedDriverLocations {
  // prevent double-add
  if (s?.byId?.[dl.driverId]) {
    return s;
  }
  return {
    byId: { ...s.byId, [dl.driverId]: dl },
    allIds: [...s.allIds, dl.driverId],
  } as NormalizedDriverLocations;
}

export function getDriverLocationsFromState(
  s: NormalizedDriverLocations
): DriverLocation[] {
  return Object.entries(s?.byId ?? []).map(([id, dl], i: number) => dl);
}

export interface GetNearestDriversResponse {
  results: SearchResult[];
}

export interface SearchResult {
  driver: DriverLocation;
  distance_meters: number;
}
