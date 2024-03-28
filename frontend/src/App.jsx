import { RouterProvider, createBrowserRouter } from "react-router-dom";

import LoginPage, { action as loginAction } from "./pages/Login";
import SignupPage, { action as signupAction } from "./pages/Signup";
import CitizensPage, {
  action as citizenAction,
  loader as citizenLoader,
} from "./pages/Citizens.jsx";

import AddPage, { action as addAction } from "./pages/AddCitizen";
import EditPage, { action as editAction } from "./pages/EditCitizen";
import ErrorPage from "./pages/Error.jsx";
import NotFound from "./pages/CatchAll.jsx";
const router = createBrowserRouter([
  {
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
    action: addAction,
    id: "add",
    errorElement: <ErrorPage></ErrorPage>,
  },
  {
    path: "/edit-citizen",
    element: <EditPage></EditPage>,
    action: editAction,
    id: "edit",
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
    path: "",
    element: <NotFound></NotFound>,
  },
]);

function App() {
  return <RouterProvider router={router} />;
}

export default App;
