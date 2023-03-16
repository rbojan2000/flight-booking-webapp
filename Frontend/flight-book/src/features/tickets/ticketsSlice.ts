import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import agent from "../../app/api/agent";

interface TicketsState {
  tickets: any[];
}

const initialState: TicketsState = {
  tickets: []
}

export const fetchTickets = createAsyncThunk<any[], void>(
  'tickets/fetchTicketsAsync',
  
  async (_, thunkAPI) => {
    try {
      
      var id = "6412e7334fd178e993697e52";
      const response = await agent.Tickets.ticketsForUser(id);
      console.log(response);
      return response;
    } catch (error: any) {
      return thunkAPI.rejectWithValue({ error: error.data });
    }
  }
);

export const ticketSlice = createSlice({
  name: "tickets",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(fetchTickets.fulfilled, (state, { payload }) => {
      state.tickets = payload;
    });
  }
});