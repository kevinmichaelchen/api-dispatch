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