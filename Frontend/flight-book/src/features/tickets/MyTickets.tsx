import { useEffect } from "react";
import { useSelector } from "react-redux";
import { fetchTickets } from "./ticketsSlice";
import { Ticket } from "../../app/models/ticket";
import { useDispatch } from "react-redux";
import { AppDispatch } from "../../app/store/configureStore";

interface Props {
    tickets: Ticket[];
}

export default function MyTickets(){
    const tickets = useSelector((state: any) => state.tickets);
    const dispatch: AppDispatch = useDispatch();
    console.log(tickets)
    
    useEffect(() => {
        dispatch(fetchTickets());
    }, []);

    return(
        <div>
            <h1> MyTickets page</h1> 
           
        </div>
    )
}