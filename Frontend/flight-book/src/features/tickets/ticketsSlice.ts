import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import agent from "../../app/api/agent";
import { Ticket } from "../../app/models/ticket";

interface TicketsState {
  tickets: Ticket[];
}

const initialState: TicketsState = {
  tickets: []
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
