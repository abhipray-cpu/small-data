import { tokenLoader } from "../util/Authentication";
import { redirect } from "react-router-dom";
export default function CitizensPage() {
  return (
    <section>
      <h1>Maa ka bhosda</h1>
    </section>
  );
}

export async function loader() {
  const token = tokenLoader();
  if (token) {
    return token;
  }
  return redirect("/login");
}
export async function action() {}
