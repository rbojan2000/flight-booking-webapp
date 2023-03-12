/* eslint-disable @typescript-eslint/no-unused-vars */
import { AppBar,  Box, List, ListItem, Switch, Toolbar, Typography } from "@mui/material";
import { NavLink } from "react-router-dom";

const midLinks = [
    {title: 'flights', path: '/flights'}
]

const rightLinks = [
    {title: 'login', path: '/login'},
    {title: 'register', path: '/register'}
]

const navStyles = {
    color: 'inherit',
    textDecoration: 'none',
    typography: 'h6',
    '&:hover':{
        color: 'grey.500'
    },
    '&.active':{
        color: 'text.secondary'
    }
}

export default function Header() {
    

    return (
        <AppBar position='static'>
            <Toolbar sx={{display: 'flex', justifyContent: 'space-between', alignItems: 'center'}}>
                
            <Box display='flex' alignItems='center'>
                    <Typography variant="h6" component={NavLink}
                        to='/'
                        sx={navStyles}
                    >
                        Flight booking
                    </Typography>
                   
                </Box>
                
                <List sx={{display: 'flex'}}>
                    {midLinks.map(({title,path}) => (
                        <ListItem
                            component={NavLink}
                            to={path}
                            key={path}
                            sx={navStyles}
                        >
                            {title.toUpperCase()}
                        </ListItem>
                    ))}
                    
                </List>
                
                <Box display='flex' alignItems='center'>
                
                <List sx={{display: 'flex'}}>
                    {rightLinks.map(({title,path}) => (
                        <ListItem
                            component={NavLink}
                            to={path}
                            key={path}
                            sx={navStyles}
                        >
                            {title.toUpperCase()}
                        </ListItem>
                    ))}
                </List>
                    
                </Box>

            </Toolbar>
        </AppBar>
    )
}