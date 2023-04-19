/* eslint-disable @typescript-eslint/no-unused-vars */

import { AppBar, Toolbar, Typography, List, ListItem } from "@mui/material";
import { Box } from "@mui/system";
import { NavLink } from "react-router-dom";
import { useSelector } from "react-redux";
import { AppDispatch, RootState } from "../../app/store/configureStore";
import { signOut } from "../../features/account/accountSlice";
import { useDispatch } from "react-redux";

const navStyles = {
  color: "inherit",
  textDecoration: "none",
  typography: "h6",
  "&:hover": {
    color: "grey.500",
  },
  "&.active": {
    color: "text.secondary",
  },
};

export default function Header() {
  const user = useSelector((state: RootState) => state.account.user);
  const dispatch = useDispatch();

  const handleLogout = () => {
    dispatch(signOut());
  };
  
  let midLinks = [
    { title: "flights", path: "/flights" },
  ];

  let rightLinks: { title: string; path: string; onClick?: () => void }[] = [
    { title: "login", path: "/login" },
    { title: "register", path: "/register" },
  ];
    

  if (user && user.roles?.includes("USER")) {
    midLinks = [
      { title: "flights", path: "/flights" },
      { title: "buy ticket", path: "/buyTicket" },
      { title: "my tickets", path: "/myTickets" },
    ];

    rightLinks = [
      { title: "log out", path: "/", onClick: handleLogout },
    ];
    
  }

if (user && user.roles?.includes("ADMIN")) {
    midLinks = [
      { title: "flights", path: "/flights" },
    ];

    rightLinks = [
      { title: "log out", path: "/", onClick: handleLogout },
    ];
    
  }

  return (
    <AppBar position="static">
      <Toolbar
        sx={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
        }}
      >
        <Box display="flex" alignItems="center">
          <Typography variant="h6" component={NavLink} to="/" sx={navStyles}>
            Flight booking
          </Typography>
        </Box>

        <List sx={{ display: "flex" }}>
          {midLinks.map(({ title, path }) => (
            <ListItem component={NavLink} to={path} key={path} sx={navStyles}>
              {title.toUpperCase()}
            </ListItem>
          ))}
        </List>

        <Box display="flex" alignItems="center">
          <List sx={{ display: "flex" }} >
            {rightLinks.map(({ title, path, onClick }) => (
              <ListItem component={NavLink} to={path} key={path} sx={navStyles} onClick={onClick}>
                {title.toUpperCase()}
              </ListItem>
            ))}
          </List>
        </Box>
      </Toolbar>
    </AppBar>
  );
}
