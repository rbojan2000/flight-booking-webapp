import { FieldValues } from "react-hook-form/dist/types";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import agent from "../../app/api/agent";
import { Flight } from "../../app/models/flight";
import { toast } from "react-toastify";

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
      const response = await agent.Flights.create(data)
        .then(() => {
          toast.success("Successful flight creation !");
        })
        .catch((error: any) => toast.error(error));
      return response;
    } catch (error: any) {
      return thunkAPI.rejectWithValue({ error: error.data });
    }
  }
);

export const removeFlight = createAsyncThunk<any, FieldValues>(
  "flights/removeFlight",

  async (data, thunkAPI) => {
    try {
      console.log(`Remove flight with ID: ${data.flightId}`);
      const response = await agent.Flights.remove(data.flightId)
        .then(() => {
          toast.success("Successful remove flight" + data.flightId + "!");
        })
        .catch((error: any) => toast.error(error));
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
