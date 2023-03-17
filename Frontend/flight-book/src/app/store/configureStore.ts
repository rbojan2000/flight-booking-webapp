import { TypedUseSelectorHook, useSelector } from 'react-redux/es/exports';
import { useDispatch } from 'react-redux/es/hooks/useDispatch';
import { configureStore } from "@reduxjs/toolkit";
import { accountSlice } from '../../features/account/accountSlice';
<<<<<<< Updated upstream
import { ticketSlice } from '../../features/tickets/ticketsSlice';
=======
>>>>>>> Stashed changes


export const store = configureStore({
    reducer: {
<<<<<<< Updated upstream
        account: accountSlice.reducer,
        ticketsSlice: ticketSlice.reducer
=======
        account: accountSlice.reducer
>>>>>>> Stashed changes
    }
})

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;