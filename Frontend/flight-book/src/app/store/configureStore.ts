import { configureStore } from "@reduxjs/toolkit";
import { useDispatch, TypedUseSelectorHook, useSelector } from "react-redux";
import { accountSlice } from "../../features/account/accountSlice";
import { flightSlice } from "../../features/flights/flightsSlice";
import { ticketSlice } from "../../features/tickets/ticketsSlice";


export const store = configureStore({
    reducer: {
        ticketsSlice: ticketSlice.reducer,
        account: accountSlice.reducer,
        flightSlice: flightSlice.reducer,
    }
})

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;