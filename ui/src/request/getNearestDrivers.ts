import { GetNearestDriversResponse, LatLng } from "../types";
import { BASE_INIT, BASE_URL } from "./base";

export default async function (
  pickupLocation: LatLng
): Promise<GetNearestDriversResponse> {
  const response = await fetch(
    `${BASE_URL}/coop.drivers.dispatch.v1beta1.DispatchService/GetNearestDrivers`,
    {
      ...BASE_INIT,
      body: JSON.stringify({ limit: 10, pickup_location: pickupLocation }),
    }
  );
  const body: GetNearestDriversResponse = await response.json();
  console.log("body", body);
  return body;
}
