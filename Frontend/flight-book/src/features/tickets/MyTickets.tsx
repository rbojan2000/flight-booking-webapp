import { useEffect } from "react";
import { useSelector } from "react-redux";
import { fetchTickets } from "./ticketsSlice";
import { useDispatch } from "react-redux";
import { AppDispatch } from "../../app/store/configureStore";
import { Ticket } from "../../app/models/ticket";
import { RootState } from "../../app/store/configureStore";
import "./styles.css";

interface Props {
    tickets: Ticket[];
}
export default function MyTickets() {
  const dispatch: AppDispatch = useDispatch();
  const tickets = useSelector((state: RootState) => state.ticketsSlice.tickets);


    useEffect(() => {
        dispatch(fetchTickets());
    }, []);
  return (
    <div className="ticket-container">
        {tickets.map(ticket => (
            <div className="ticket-card" key={ticket.ID}>
                <h2>{ticket.Flight.Departure.City} - {ticket.Flight.Arrival.City}</h2>
                <p>Price: ${ticket.Flight.Price}</p>
                <p>Date and time: {ticket.Flight.Date}</p>

                {/* <p>Passenger: {ticket.User.Name} {ticket.User.Surname}</p> */}
            </div>
        ))}
    </div>
  );
}