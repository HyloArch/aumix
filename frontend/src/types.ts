import type React from "react";

export type PageType = {
  id: string;
  name: string;
  page: () => React.JSX.Element;
};

export type MessageOp = "GET" | "SET" | "GET_OSC" | "SET_OSC";

export type Message = {
  op: MessageOp;
  key: string;
  value?: any;
};

export type MeterMapping = {
  name: string;
  fader: {
    type: string;
    index: number;
  };
  meter: {
    typeId: number;
    index: number;
  };
};

export type Sample = {
  name: string;
  file: string;
};

export type Show = {
  id: number;
  name: string;
  scenes: ShowScene[];
};

export type ShowScene = {
  id: number;
  name: string;
  sceneId: number;
  movement: number;
  measure: number;
  samples?: Sample[];
};
