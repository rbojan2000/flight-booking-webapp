import { FieldValues } from "react-hook-form/dist/types";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import agent from "../../app/api/agent";
import { Flight } from "../../app/models/flight";

interface FlightsState {
  flights: Flight[];
  departure: string;
  arrival: string;
  price: string;
  date: string;
}

const initialState: FlightsState = {
  flights: [],
  departure: "",
  arrival: "",
  price: "",
  date: "",
};

export const fetchFlights = createAsyncThunk<any[], void>(
  "flights/fetchFlightsAsync",

  async (_, thunkAPI) => {
    try {
      const response = await agent.Flights.flights();
      return response;
    } catch (error: any) {
      return thunkAPI.rejectWithValue({ error: error.data });
    }
  }
);

export const createFlight = createAsyncThunk<any, FieldValues>(
  "flights/createFlight",

  async (data, thunkAPI) => {
    try {
      console.log(data);
      const response = await agent.Flights.create(data);
      return response;
    } catch (error: any) {
      return thunkAPI.rejectWithValue({ error: error.data });
    }
  }
);

export const flightSlice = createSlice({
  name: "flights",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(fetchFlights.fulfilled, (state, { payload }) => {
      state.flights = payload;
    });
  },
});
