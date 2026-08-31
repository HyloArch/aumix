import type { MeterMapping } from "./types";

// export const HOST = window.location.host
export const HOST = "localhost:8080";

export const MeterMappings: { [s: string]: MeterMapping } = {
  Ch1: {
    name: "Ch1",
    fader: {
      type: "ch",
      index: 1,
    },
    meter: {
      typeId: 0,
      index: 0,
    },
  },
  Ch2: {
    name: "Ch2",
    fader: {
      type: "ch",
      index: 2,
    },
    meter: {
      typeId: 0,
      index: 1,
    },
  },
  Ch3: {
    name: "Ch3",
    fader: {
      type: "ch",
      index: 3,
    },
    meter: {
      typeId: 0,
      index: 2,
    },
  },
  Ch4: {
    name: "Ch4",
    fader: {
      type: "ch",
      index: 4,
    },
    meter: {
      typeId: 0,
      index: 3,
    },
  },
  Ch5: {
    name: "Ch5",
    fader: {
      type: "ch",
      index: 5,
    },
    meter: {
      typeId: 0,
      index: 4,
    },
  },
  Ch6: {
    name: "Ch6",
    fader: {
      type: "ch",
      index: 6,
    },
    meter: {
      typeId: 0,
      index: 5,
    },
  },
  Ch7: {
    name: "Ch7",
    fader: {
      type: "ch",
      index: 7,
    },
    meter: {
      typeId: 0,
      index: 6,
    },
  },
  Ch8: {
    name: "Ch8",
    fader: {
      type: "ch",
      index: 8,
    },
    meter: {
      typeId: 0,
      index: 7,
    },
  },
} as const;
