import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { FieldValues } from "react-hook-form/dist/types";
import agent from "../../app/api/agent";
import { Ticket } from "../../app/models/ticket";

import { Flight } from "../../app/models/flight";
import { toast } from "react-toastify";

interface TicketsState {
  tickets: Ticket[];
  flights: Flight[];
}

const initialState: TicketsState = {
  tickets: [],
  flights: []
}

export const fetchTickets = createAsyncThunk<any[], void>(
  'tickets/fetchTicketsAsync',
  
  async (_, thunkAPI) => {
    try {
      
      var id = "6413607fc2fac0c7689d944b";
      const response = await agent.Tickets.ticketsForUser(id);
      console.log(response);
      return response;
    } catch (error: any) {
      return thunkAPI.rejectWithValue({ error: error.data });
    }
  }
);


export const fetchFlights = createAsyncThunk<any[], void>(
  "tickets/fetchFlightsAsync",
  async (_, thunkAPI) => {
    try {
      const response = await agent.Tickets.flights();
      return response;
    } catch (error: any) {
      return thunkAPI.rejectWithValue({ error: error.data });
    }
  }
);


export const createTicket = createAsyncThunk<any, FieldValues>(
  "/buyTicket",
  async (data, thunkAPI) => {
    try {
      var id = "6413607fc2fac0c7689d944b";
      let buyTicketDTO = {flightID: data.selectedFlightInfo.ID, userID: id, numberOfTickets:data.numberOfTickets}   
      await agent.Tickets.create(buyTicketDTO);
   
      return true;
    } catch (error: any) {
      return thunkAPI.rejectWithValue({ error: error.data });
    }
  }
)

export const ticketSlice = createSlice({
  name: "tickets",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(fetchTickets.fulfilled, (state, { payload }) => {
      state.tickets = payload;
    });
     builder.addCase(fetchFlights.fulfilled, (state, { payload }) => {
      state.flights = payload;
    });

    builder.addCase(createTicket.rejected, (state, { payload }) => {
      const errorMessage = (payload as { error: string }).error;
      toast.error(errorMessage)
    });
    builder.addCase(createTicket.fulfilled, (state, { payload }) => {
      toast.success("Congrats! You bought tickets.")
    });
    
  }
});