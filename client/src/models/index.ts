import { ReactElement } from "react";

export interface ComponentWithChildren {
  children: ReactElement | ReactElement[];
}

export type IAddInterface = {
  person: string;
  productname: string;
  purchase: number;
};
export type IUpdateOnlyInterface = {
  id: string;
  productname: string;
  purchase: number;
};

export type IUpdateAllInterface = {
  id: string;
  person: string;
  productname: string;
  purchase: number;
};

type Date = {
  day: string;
  month: string;
  year: string;
  time: string;
};

export type IncExpModels = [
  {
    id?: string;
    category?: string;
    person: string;
    productname: string;
    purchase: number;
    date?: Date;
  }
];

export type IncExpModel = 
  {
    id?: string;
    category?: string;
    person: string;
    productname: string;
    purchase: number;
    date?: Date;
  }






