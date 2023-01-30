import { IconButton, Stack, Typography } from "@mui/material";
import { Add as AddIcon, Remove as RemoveIcon } from "@mui/icons-material";
import { Dispatch, SetStateAction } from "react";

interface ResolutionControlProps {
  resolution: number;
  setResolution: Dispatch<SetStateAction<number>>;
}

const MAX_RESOLUTION = 10;
const MIN_RESOLUTION = 7;

export default function ResolutionControl(props: ResolutionControlProps) {
  const { resolution, setResolution } = props;
  return (
    <Stack direction="row">
      <IconButton
        aria-label="remove"
        onClick={() => {
          if (resolution > MIN_RESOLUTION) setResolution(resolution - 1);
        }}
      >
        <RemoveIcon />
      </IconButton>
      <Typography variant="h3">{resolution}</Typography>
      <IconButton
        aria-label="add"
        onClick={() => {
          if (resolution < MAX_RESOLUTION) setResolution(resolution + 1);
        }}
      >
        <AddIcon />
      </IconButton>
    </Stack>
  );
}
