import { BASE_INIT, BASE_URL } from "./base";
import { DriverLocation } from "../types";

export default async function (
  driverLocations: DriverLocation[]
): Promise<void> {
  const now = new Date().toISOString();
  await fetch(
    `${BASE_URL}/coop.drivers.dispatch.v1beta1.DispatchService/UpdateDriverLocations`,
    {
      ...BASE_INIT,
      body: JSON.stringify({
        locations: driverLocations.map((dl) => ({
          ...dl,
          mostRecentHeartbeat: now,
        })),
      }),
    }
  );
}
