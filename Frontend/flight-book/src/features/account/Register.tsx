import Avatar from "@mui/material/Avatar";
import TextField from "@mui/material/TextField";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
import { Paper } from "@mui/material";
import { useForm } from "react-hook-form";
import { Link, useNavigate } from "react-router-dom";
import agent from "../../app/api/agent";
import { toast } from "react-toastify";
import { LoadingButton } from "@mui/lab";

export default function Register() {
  const navigate = useNavigate();
  const {
    register,
    handleSubmit,
    setError,
    formState: { isSubmitting, errors, isValid },
  } = useForm({
    mode: "onTouched",
  });

  function handleApiErorrs(error: any) {
    if (error) {
      if (error.data.includes("Password")) {
        setError("password", { message: error.data });
      } else if (error.data.includes("Email")) {
        setError("email", { message: error.data });
      } else if (error.data.includes("Username")) {
        setError("username", { message: error.data });
      }
    }
  }

  return (
    <Container
      component={Paper}
      maxWidth="sm"
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        p: 4,
      }}
    >
      <Avatar sx={{ m: 1, bgcolor: "secondary.main" }}>
        <LockOutlinedIcon />
      </Avatar>
      <Typography component="h1" variant="h5">
        Register
      </Typography>
      <Box
        component="form"
        onSubmit={handleSubmit((data) => {
          const { name, surname, email, password, country, city } = data;
          const newData = {
            name,
            surname,
            email,
            password,
            address: {
              country,
              city,
            },
          };
          agent.Account.register(newData)
            .then(() => {
              toast.success("Registration successful - you can now login");
              navigate("/login");
            })
            .catch((error: any) => handleApiErorrs(error));
        })}
        noValidate
        sx={{ mt: 1 }}
      >
        <TextField
          margin="normal"
          fullWidth
          label="Name"
          autoFocus
          {...register("name", { required: "Name is required" })}
          error={!!errors.name}
          helperText={errors?.name?.message as string}
        />
        <TextField
          margin="normal"
          fullWidth
          label="Surname"
          autoFocus
          {...register("surname", { required: "Surname is required" })}
          error={!!errors.surname}
          helperText={errors?.surname?.message as string}
        />
        <TextField
          margin="normal"
          fullWidth
          label="Email"
          {...register("email", {
            required: "Email is required",
            pattern: {
              value: /^\w+[\w-.]*@\w+((-\w+)|(\w*))\.[a-z]{2,3}$/,
              message: "Not a valid email address",
            },
          })}
          error={!!errors.email}
          helperText={errors?.email?.message as string}
        />
        <TextField
          margin="normal"
          fullWidth
          label="Password"
          type="password"
          {...register("password", {
            required: "Password is required",
            pattern: {
              value:
                /(?=^.{6,10}$)(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&amp;*()_+}{&quot;:;'?/&gt;.&lt;,])(?!.*\s).*$/,
              message: "Password does not meet complexity requirements",
            },
          })}
          error={!!errors.password}
          helperText={errors?.password?.message as string}
        />
        <TextField
          margin="normal"
          fullWidth
          label="Country"
          autoFocus
          {...register("country", { required: "Country is required" })}
          error={!!errors.country}
          helperText={errors?.country?.message as string}
        />
        <TextField
          margin="normal"
          fullWidth
          label="City"
          autoFocus
          {...register("city", { required: "City is required" })}
          error={!!errors.city}
          helperText={errors?.city?.message as string}
        />
        <LoadingButton
          disabled={!isValid}
          loading={isSubmitting}
          type="submit"
          fullWidth
          variant="contained"
          sx={{ mt: 3, mb: 2 }}
        >
          Register
        </LoadingButton>
        <Grid container>
          <Grid item>
            <Link to="/login">{"Already have an account? Sign In"}</Link>
          </Grid>
        </Grid>
      </Box>
    </Container>
  );
}
