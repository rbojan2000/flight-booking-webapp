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
        PaperProps={{
        className: "ticket-dialog",
        }}
      >

        <DialogTitle className="ticket-dialog-title">Select Number of Tickets</DialogTitle>
        <DialogContent className="ticket-dialog-content">
          <Typography>
          <b><FaCalendarAlt /> Date:</b>{" "}
            {new Date(
              flights.find((flight) => flight.ID === selectedFlight)?.Date ?? ""
            ).toLocaleString()}
            <br />
            <b> <FaPlaneDeparture/>Departure:</b>{" "}
            {flights.find((flight) => flight.ID === selectedFlight)?.Departure.City},{" "}
            {flights.find((flight) => flight.ID === selectedFlight)?.Departure.Country}
            <br />
            <b> <FaPlaneArrival/> Arrival:</b>{" "}
            {flights.find((flight) => flight.ID === selectedFlight)?.Arrival.City},{" "}
            {flights.find((flight) => flight.ID === selectedFlight)?.Arrival.Country}
            <br />
            <b> <GiPriceTag/> Price per ticket:</b>{" "}
            {flights.find((flight) => flight.ID === selectedFlight)?.Price} $
          </Typography>
          <br />
          <TextField
            label="Number of Tickets"
            type="number"
            value={numberOfTickets}
            onChange={(event) => setNumberOfTickets(Number(event.target.value))}
            
            InputProps={{ inputProps: { 
              min: 1,
              max: 10
            } }}
            fullWidth
          />
          <br />
          <Typography variant="h6" className="ticket-dialog-price">
            {selectedFlight &&
            flights.find(
              (flight) =>
                flight.ID === selectedFlight && flight.Price !== undefined
            ) ? (
              flights.find((flight) => flight.ID === selectedFlight)?.Price! *
              numberOfTickets
            ) : (
              0
            )}{" "}
            $
          </Typography>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleDialogClose}>Cancel</Button>
          <Button onClick={handleDialogSubmit} color="primary">
            Buy
          </Button>
        </DialogActions>
      </Dialog>

  </Container>
    )
}