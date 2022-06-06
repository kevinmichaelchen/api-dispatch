import React, { MouseEventHandler } from "react";
import {
  Box,
  Button,
  Stack,
  List,
  ListItem,
  Typography,
  Divider,
  Avatar,
} from "@mui/material";
import { DriverLocation } from "../types";

function stringToColor(string: string) {
  let hash = 0;
  let i;

  /* eslint-disable no-bitwise */
  for (i = 0; i < string.length; i += 1) {
    hash = string.charCodeAt(i) + ((hash << 5) - hash);
  }

  let color = "#";

  for (i = 0; i < 3; i += 1) {
    const value = (hash >> (i * 8)) & 0xff;
    color += `00${value.toString(16)}`.slice(-2);
  }
  /* eslint-enable no-bitwise */

  return color;
}

function stringAvatar(name: string) {
  if (!name) {
    return {};
  }
  const split = name.split("-");
  return {
    sx: {
      bgcolor: stringToColor(name),
    },
    children:
      split.length > 1
        ? `${split[0][0]}${split[split.length - 1][0]}`
        : `${split[0][0]}`,
  };
}

interface ItemProps {
  driverLocation: DriverLocation;
  buildHandleMouseOver?: (driverId: string) => MouseEventHandler;
  buildHandleMouseOut?: (driverId: string) => MouseEventHandler;
}

function Item(props: ItemProps) {
  const {
    driverLocation: dl,
    buildHandleMouseOver,
    buildHandleMouseOut,
  } = props;
  const { driverId } = dl;
  return (
    <ListItem
      onMouseOver={buildHandleMouseOver && buildHandleMouseOver(driverId)}
      onMouseOut={buildHandleMouseOut && buildHandleMouseOut(driverId)}
    >
      <Stack direction="row" spacing={2} alignItems="center">
        <Avatar alt={`Driver: ${driverId}`} {...stringAvatar(driverId)} />
        <Typography>{driverId}</Typography>
      </Stack>
    </ListItem>
  );
}

interface DriverListProps {
  driverLocations: DriverLocation[];
  onUpload?: MouseEventHandler;
  /**
   * Whether the list is showing nearby ranked drivers (Query Mode) or cached drivers that will be created (Mutate Mode)
   */
  queryMode: boolean;
  buildHandleMouseOver?: (driverId: string) => MouseEventHandler;
  buildHandleMouseOut?: (driverId: string) => MouseEventHandler;
}

export default function DriverList(props: DriverListProps) {
  const {
    driverLocations,
    onUpload: handleClick,
    queryMode,
    buildHandleMouseOver,
    buildHandleMouseOut,
  } = props;
  return (
    <Stack spacing={1} px="20px" py="20px">
      <Typography variant="h5" component="div">
        {queryMode ? "Drivers List" : "New Drivers List"}
      </Typography>
      <Divider />
      <List>
        {driverLocations.map((p: DriverLocation, i: number) => (
          <Item
            key={i}
            driverLocation={p}
            buildHandleMouseOver={buildHandleMouseOver}
            buildHandleMouseOut={buildHandleMouseOut}
          />
        ))}
      </List>
      {!queryMode && (
        <Stack justifyContent="center" alignItems="center">
          <Button
            onClick={handleClick}
            sx={{ maxWidth: "100px" }}
            variant="outlined"
          >
            Submit
          </Button>
        </Stack>
      )}
    </Stack>
  );
}
