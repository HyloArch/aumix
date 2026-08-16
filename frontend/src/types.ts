import type React from "react";

export type PageType = {
  id: string;
  name: string;
  page: React.JSX.Element;
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
