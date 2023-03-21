// import { Close } from "@mui/icons-material";
// import {
//   TableContainer,
//   Paper,
//   Table,
//   TableHead,
//   TableRow,
//   TableCell,
//   TableBody,
//   Button,
//   IconButton,
//   Modal,
//   TextField,
//   Typography,
// } from "@mui/material/";
// import moment from "moment";
// import { useEffect, useState } from "react";
// import { useDispatch, useSelector } from "react-redux";
// import { AppDispatch, RootState } from "../../app/store/configureStore";
// import { fetchFlights } from "./flightsSlice";
// import { makeStyles } from "@material-ui/core/styles";
// import { createFlight } from "./flightsSlice";

// const useStyles = makeStyles((theme: any) => ({
//   modal: {
//     display: "flex",
//     alignItems: "center",
//     justifyContent: "center",
//   },
//   paper: {
//     backgroundColor: theme.palette.background.paper,
//     boxShadow: theme.shadows[5],
//     padding: theme.spacing(2, 4, 3),
//     minWidth: 300,
//     borderRadius: 5,
//   },
//   closeButton: {
//     position: "absolute",
//     right: theme.spacing(1),
//     top: theme.spacing(1),
//     color: theme.palette.grey[500],
//   },
//   modalBackdrop: {
//     backgroundColor: "rgba(0, 0, 0, 0.5)",
//   },
// }));

export default function FlightsPage() {
//   const dispatch: AppDispatch = useDispatch();
//   const flights = useSelector((state: RootState) => state.flightSlice.flights);
//   const [date, setDate] = useState("");
//   const [time, setTime] = useState("");
//   const [TicketNum, setTicketNum] = useState("");
//   const [Price, setPrice] = useState("");
//   const [ArrivalCountry, setArrivalCountry] = useState("");
//   const [ArrivalCity, setArrivalCity] = useState("");
//   const [DepartureCountry, setDepartureCountry] = useState("");
//   const [DepartureCity, setDepartureCity] = useState("");

//   const [open, setOpen] = useState(false);

//   const handleOpen = () => {
//     setOpen(true);
//   };
//   const handleClose = () => {
//     setOpen(false);
//   };
//   const classes = useStyles();

//   useEffect(() => {
//     dispatch(fetchFlights());
//   }, [dispatch]);

//   const handleSubmit = () => {
//     const DateAndTime = date + ", " + time;
//     dispatch(
//       createFlight({
//         ArrivalCountry,
//         ArrivalCity,
//         DepartureCountry,
//         DepartureCity,
//         DateAndTime,
//         TicketNum,
//         Price,
//       })
//     );
//   };

  return (
//     <>
<div></div>
//       <TableContainer component={Paper}>
//         <Table>
//           <TableHead>
//             <TableRow>
//               <TableCell>Departure</TableCell>
//               <TableCell>ArrivalCountry</TableCell>
//               <TableCell>Date and Time</TableCell>
//               <TableCell>Price</TableCell>
//               <TableCell>Remaining tickets</TableCell>
//             </TableRow>
//           </TableHead>
//           <TableBody>
//             {flights.map((flight) => (
//               <TableRow key={flight.ID}>
//                 <TableCell>
//                   {flight.Departure.Country}, {flight.Departure.City}
//                 </TableCell>
//                 <TableCell>
//                   {flight.Arrival.Country}, {flight.Arrival.City}
//                 </TableCell>
//                 <TableCell>
//                   {moment.utc(flight.Date).utc().format("DD.MM.YYYY, HH:mm")}
//                 </TableCell>
//                 <TableCell>{flight.Price}</TableCell>
//                 <TableCell>
//                   {flight.PassengerCount}/{flight.Capacity}
//                 </TableCell>
//               </TableRow>
//             ))}
//           </TableBody>
//         </Table>
//       </TableContainer>

//       <>
//         <Button variant="contained" color="primary" onClick={handleOpen}>
//           Create
//         </Button>
//         <Modal
//           open={open}
//           onClose={handleClose}
//           onSubmit={handleSubmit}
//           className={classes.modal}
//           BackdropProps={{
//             classes: {
//               root: classes.modalBackdrop,
//             },
//           }}
//         >
//           <div className={classes.paper}>
//             <Typography variant="h5">Create a flight</Typography>
//             <IconButton className={classes.closeButton} onClick={handleClose}>
//               <Close />
//             </IconButton>
//             <form>
//               <TextField
//                 label="Departure Country"
//                 type="text"
//                 value={DepartureCountry}
//                 onChange={(event) => setDepartureCountry(event.target.value)}
//                 variant="outlined"
//                 margin="dense"
//                 fullWidth
//               />
//               <TextField
//                 label="Departure City"
//                 type="text"
//                 value={DepartureCity}
//                 onChange={(event) => setDepartureCity(event.target.value)}
//                 variant="outlined"
//                 margin="dense"
//                 fullWidth
//               />
//               <TextField
//                 label="Arrival Country"
//                 type="text"
//                 value={ArrivalCountry}
//                 onChange={(event) => setArrivalCountry(event.target.value)}
//                 variant="outlined"
//                 margin="dense"
//                 fullWidth
//               />
//               <TextField
//                 label="Arrival City"
//                 type="text"
//                 value={ArrivalCity}
//                 onChange={(event) => setArrivalCity(event.target.value)}
//                 variant="outlined"
//                 margin="dense"
//                 fullWidth
//               />
//               <TextField
//                 type="date"
//                 value={date}
//                 onChange={(event) => setDate(event.target.value)}
//                 variant="outlined"
//                 margin="dense"
//                 fullWidth
//               />
//               <TextField
//                 type="time"
//                 value={time}
//                 onChange={(event) => setTime(event.target.value)}
//                 variant="outlined"
//                 margin="dense"
//                 fullWidth
//               />
//               <TextField
//                 label="Capacity"
//                 type="text"
//                 value={TicketNum}
//                 onChange={(event) => setTicketNum(event.target.value)}
//                 variant="outlined"
//                 margin="dense"
//                 fullWidth
//               />
//               <TextField
//                 label="Price"
//                 type="text"
//                 value={Price}
//                 onChange={(event) => setPrice(event.target.value)}
//                 variant="outlined"
//                 margin="dense"
//                 fullWidth
//               />
//               <Button
//                 type="button"
//                 onClick={handleSubmit}
//                 variant="contained"
//                 color="primary"
//               >
//                 Submit
//               </Button>
//             </form>
//           </div>
//         </Modal>
//       </>
//     </>
  );
}
