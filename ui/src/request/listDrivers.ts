import { DriverLocation } from "../types";
import { BASE_INIT, BASE_URL } from "./base";

export default async function (): Promise<DriverLocation[]> {
  const response = await fetch(
    `${BASE_URL}/coop.drivers.dispatch.v1beta1.DispatchService/ListDrivers`,
    {
      ...BASE_INIT,
      body: JSON.stringify({ page_size: 1000 }),
    }
  );
  const body = await response.json();
  console.log("body", body);
  return body.driverLocations;
}
