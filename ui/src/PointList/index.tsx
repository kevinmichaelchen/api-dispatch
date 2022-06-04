import React from "react";
import {
  Box,
  Button,
  Stack,
  List,
  ListItem,
  Typography,
  Divider,
} from "@mui/material";
import { Point } from "../types";

function PointListItem({ point }: { point: Point }) {
  return <ListItem>{point.lat}</ListItem>;
}

interface PointListProps {
  points: Point[];
}

export default function PointList(props: PointListProps) {
  const { points } = props;
  return (
    <Stack spacing={1} px="20px" py="20px">
      <Typography variant="h5" component="div">
        Point List
      </Typography>
      <Divider />
      <List>
        {points.map((p: Point, i: number) => (
          <PointListItem point={p} key={i} />
        ))}
      </List>
      <Stack justifyContent="center" alignItems="center">
        <Button sx={{ maxWidth: "100px" }} variant="outlined">
          Submit
        </Button>
      </Stack>
    </Stack>
  );
}
