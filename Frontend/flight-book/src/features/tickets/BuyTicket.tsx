import { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { fetchFlights } from "./ticketsSlice";
import { AppDispatch, RootState } from "../../app/store/configureStore";
import { useState } from "react";
import { FaCalendarAlt } from "react-icons/fa";
import {FaPlaneDeparture, FaPlaneArrival} from "react-icons/fa";
import { createTicket } from "./ticketsSlice";
import {GiPriceTag} from "react-icons/gi";
import {
  Container,
  Typography,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
} from "@mui/material";
import "./styles.css";



export default function BuyTicket() {
  const dispatch: AppDispatch = useDispatch();
  const flights = useSelector((state: RootState) => state.ticketsSlice.flights);

  const [openDialog, setOpenDialog] = useState(false);
  const [selectedFlight, setSelectedFlight] = useState("");
  const [numberOfTickets, setNumberOfTickets] = useState(1);

  const handleBuyTicket = (flightId: string) => {
    setSelectedFlight(flightId);
    setOpenDialog(true);
    console.log(`Buy ticket for flight ID: ${flightId}`);
  };

  const handleDialogClose = () => {
    // Reset the state when the dialog is closed
    setSelectedFlight("");
    setNumberOfTickets(1);
    setOpenDialog(false);
  };

  const handleDialogSubmit = () => {
    const selectedFlightInfo = flights.find((flight) => flight.ID === selectedFlight);
    if (selectedFlightInfo) {
      // Dispatch the buyTicket action with the selected flight ID and number of tickets
     dispatch(createTicket({ selectedFlightInfo: selectedFlightInfo, numberOfTickets }));
    }
    handleDialogClose();
  };

  useEffect(() => {
    dispatch(fetchFlights());
  }, []);

  return (
    
    <Container>
      <h1>Buy Ticket page</h1>
      <Table className="table">
        <TableHead>
          <TableRow>
            <TableCell>Date</TableCell>
            <TableCell>Departure</TableCell>
            <TableCell>Arrival</TableCell>
            <TableCell>Price</TableCell>
            <TableCell>Busyness</TableCell>
              <TableCell></TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {flights.map((flight) => (
             <TableRow key={flight.ID} style={{ cursor: "pointer" }}>
                <TableCell>  {new Date(flight.Date).toLocaleString("en-US", {
                    year: "numeric",
                    month: "short",
                    day: "numeric",
                    hour: "numeric",
                    minute: "numeric",
                    second: "numeric",
                    timeZoneName: "shortGeneric",
                })}</TableCell>
                <TableCell>{flight.Departure.Country}, {flight.Departure.City}</TableCell>
                <TableCell>{flight.Arrival.Country},{flight.Arrival.City}</TableCell>
                <TableCell>{flight.Price} $</TableCell>
                <TableCell>{flight.PassengerCount}/{flight.Capacity}</TableCell>
                
                <TableCell>
                  <Button
                    variant="contained"
                    color="primary"
                    onClick={() => handleBuyTicket(flight.ID)}
                  >
                    Buy
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>

        {/* Dialog box to select the number of tickets */}
        <Dialog
          open={openDialog}
          onClose={handleDialogClose}
          sx={{
            "& .MuiDialog-paper": {
              width: "50%",
              height: "60%",
            },
          }}
        >
          <div className="my-dialog-container">
            <div className="my-dialog-content">
              <h2>Select Number of Tickets</h2>
              <p><strong><FaCalendarAlt /> Date:</strong> {new Date(flights.find((flight) => flight.ID === selectedFlight)?.Date ?? "").toLocaleString()}</p>
              <p><strong><FaPlaneDeparture/> Departure:</strong> {flights.find((flight) => flight.ID === selectedFlight)?.Departure.City}, {flights.find((flight) => flight.ID === selectedFlight)?.Departure.Country}</p>
              <p><strong><FaPlaneArrival/> Arrival:</strong> {flights.find((flight) => flight.ID === selectedFlight)?.Arrival.City}, {flights.find((flight) => flight.ID === selectedFlight)?.Arrival.Country}</p>
              <p><strong><GiPriceTag/> Price per ticket:</strong> {flights.find((flight) => flight.ID === selectedFlight)?.Price} $</p>
              <label htmlFor="numberOfTickets">Number of Tickets:</label>
              <input type="number" id="numberOfTickets" name="numberOfTickets" min="1" max="10" value={numberOfTickets} onChange={(event) => setNumberOfTickets(Number(event.target.value))} />
              <p className="ticket-dialog-price">{selectedFlight && flights.find((flight) => flight.ID === selectedFlight && flight.Price !== undefined) ? flights.find((flight) => flight.ID === selectedFlight)?.Price! * numberOfTickets : 0} $</p>
            </div>
            <div className="my-dialog-buttons">
              <button onClick={handleDialogClose}>Cancel</button>
              <button onClick={handleDialogSubmit}>Buy</button>
            </div>
          </div>
        </Dialog>
      </Container>
      
    )
}