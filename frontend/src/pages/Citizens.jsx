import { useEffect } from "react";

export default function CitizensPage() {
  useEffect(() => {
    console.log("maa ka bhosda");
  }, []);
  return (
    <section>
      <h1>Maa ka bhosda</h1>
    </section>
  );
}

export async function loader() {}
export async function action() {}
