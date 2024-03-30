import { RouterProvider, createBrowserRouter } from "react-router-dom";

import LoginPage, { action as loginAction } from "./pages/Login";
import SignupPage, { action as signupAction } from "./pages/Signup";

import CitizensPage, {
  action as citizenAction,
  loader as citizenLoader,
} from "./pages/Citizens.jsx";

import DeleteCitizen, {
  loader as deleteCitizenLoader,
} from "./pages/DeleteCitizen.jsx";
import AddPage from "./pages/AddCitizen";
import EditPage, { loader as editCitizenLoader } from "./pages/EditCitizen";
import ErrorPage from "./pages/Error.jsx";
import NotFound from "./pages/CatchAll.jsx";
import { action as formAction } from "./components/CitizenForm.jsx";
const router = createBrowserRouter([
  {
    path: "/",
    element: <CitizensPage />,
    errorElement: <ErrorPage />,
    action: citizenAction,
    loader: citizenLoader,
    id: "home",
    index: true,
  },
  {
    path: "/add-citizen",
    element: <AddPage></AddPage>,
    id: "add",
    action: formAction,
    errorElement: <ErrorPage></ErrorPage>,
  },
  {
    path: "/edit-citizen/:id",
    element: <EditPage></EditPage>,
    id: "edit",
    action: formAction,
    loader: editCitizenLoader,
    errorElement: <ErrorPage></ErrorPage>,
  },
  {
    path: "/delete-citizen/:id",
    element: <DeleteCitizen />,
    id: "delete",
    loader: deleteCitizenLoader,
    errorElement: <ErrorPage></ErrorPage>,
  },
  {
    path: "/login",
    element: <LoginPage></LoginPage>,
    action: loginAction,
    id: "login",
    errorElement: <ErrorPage></ErrorPage>,
  },
  {
    path: "/signup",
    element: <SignupPage></SignupPage>,
    action: signupAction,
    id: "signup",
    errorElement: <ErrorPage></ErrorPage>,
  },
  {
    path: "*",
    element: <NotFound></NotFound>,
  },
]);

function App() {
  return <RouterProvider router={router} />;
}

export default App;
