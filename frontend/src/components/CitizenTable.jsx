import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { getToken } from "../util/Authentication";
export default function CitizenTable({ citizens }) {
  const [table, setTable] = useState(citizens.data);
  const [currentPage, setCurrentPage] = useState(1);
  const [range, setRange] = useState(1);
  const navigate = useNavigate();

  useEffect(() => {
    // using flag to check the status of this comp and prevent race condition
    let mounted = true;
    const fetchData = async () => {
      const token = getToken();
      if (!token || token === "EXPIRED") {
        navigate("/login");
      }
      const response = await fetch(
        `http://localhost:8080/citizens/${currentPage}`,
        {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            Authorization: token,
          },
        }
      );
      if (!response.ok) {
        return;
      }
      if (mounted) {
        let data = await response.json();
        if (!data.data) {
          setTable([]);
        } else setTable(data.data);
      }
    };
    fetchData();
    // Cleanup function to cancel any pending async operations if component unmounts
    return () => {
      mounted = false;
    };
  }, [currentPage, navigate]); // Placeholder useEffect, not currently used

  const nextPage = () => {
    const nextPage = currentPage + 1;
    const nextRangePage = range + 9;

    if (nextPage > nextRangePage) {
      setRange((prevRange) => prevRange + 10);
    }
    setCurrentPage(nextPage);
  };

  const prevPage = () => {
    const prevPage = currentPage - 1;

    if (prevPage < range) {
      if (range > 1) {
        setRange((prevRange) => prevRange - 10);
      }
    }
    setCurrentPage(prevPage);
  };

  const setPage = (value) => {
    setCurrentPage(value);
  };
  return (
    <section className="w-full overflow-auto">
      {table.length === 0 && (
        <h3 className="text-2xl text-gray-700 font-medium font-serif text-center tracking-wide">
          No data to load
        </h3>
      )}
      {table.length > 0 && (
        <>
          <table className="min-w-full divide-y divide-gray-200">
            <thead className="bg-gray-200">
              <tr>
                <th className="text-center px-4 py-2 text-md font-medium text-gray-500 uppercase tracking-wider w-fit">
                  First Name
                </th>
                <th className="text-center px-4 py-2 text-md font-medium text-gray-500 uppercase tracking-wider w-fit">
                  Last Name
                </th>
                <th className="text-center px-4 py-2 text-md font-medium text-gray-500 uppercase tracking-wider w-fit">
                  DOB
                </th>
                <th className="text-center px-4 py-2 text-md font-medium text-gray-500 uppercase tracking-wider w-fit">
                  Gender
                </th>
                <th className="text-center px-4 py-2 text-md font-medium text-gray-500 uppercase tracking-wider w-fit">
                  City
                </th>
                <th className="text-center px-4 py-2 text-md font-medium text-gray-500 uppercase tracking-wider w-fit">
                  PinCode
                </th>
                <th className="text-center px-4 py-2 text-md font-medium text-gray-500 uppercase tracking-wider w-fit">
                  State
                </th>
                <th className="text-center px-4 py-2 text-md font-medium text-gray-500 uppercase tracking-wider w-fit">
                  Address
                </th>
                <th className="text-center px-4 py-2 text-md font-medium text-gray-500 uppercase tracking-wider w-fit">
                  Delete
                </th>
                <th className="text-center px-4 py-2 text-md font-medium text-gray-500 uppercase tracking-wider w-fit">
                  Edit
                </th>
              </tr>
            </thead>
            <tbody>
              {table.map((citizen, index) => (
                <tr
                  key={index}
                  className={index % 2 === 0 ? "bg-gray-200" : "bg-gray-400"}
                >
                  <td className="text-center px-4 py-2 whitespace-nowrap text-sm font-medium text-gray-800">
                    {citizen.FirstName}
                  </td>
                  <td className="text-center px-4 py-2 whitespace-nowrap text-md text-gray-800">
                    {citizen.LastName}
                  </td>
                  <td className="text-center px-4 py-2 whitespace-nowrap text-md text-gray-800">
                    {citizen.DOB}
                  </td>
                  <td className="text-center px-4 py-2 whitespace-nowrap text-md text-gray-800">
                    {citizen.Gender}
                  </td>
                  <td className="text-center px-4 py-2 whitespace-nowrap text-md text-gray-800">
                    {citizen.City}
                  </td>
                  <td className="text-center px-4 py-2 whitespace-nowrap text-md text-gray-800">
                    {citizen.PinCode}
                  </td>
                  <td className="text-center px-4 py-2 whitespace-nowrap text-md text-gray-800">
                    {citizen.State}
                  </td>
                  <td className="text-center px-4 py-3 whitespace-nowrap text-md text-gray-800">
                    {citizen.Address}
                  </td>
                  <td
                    className="text-center px-4 py-3 whitespace-nowrap text-md text-white bg-red-600"
                    onClick={() => {
                      navigate(`delete-citizen/${citizen.ID}`);
                    }}
                  >
                    Delete
                  </td>
                  <td
                    className="text-center px-4 py-3 whitespace-nowrap text-md text-white bg-orange-600"
                    onClick={() => {
                      navigate(`edit-citizen/${citizen.ID}`);
                    }}
                  >
                    Edit
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </>
      )}

      <div className="bottom-15 mt-5 w-full flex absolute justify-center items-center">
        <button onClick={prevPage} className="text-s">
          Prev
        </button>
        <ul className="flex gap-1">
          {Array.from({ length: 10 }, (_, i) => {
            const pageNumber = range + i;
            const isActive = pageNumber === currentPage;
            return (
              <li key={i} className="mx-[0.25vw]">
                <button
                  className={`px-1 py-1 rounded ${
                    isActive ? "bg-gray-800 text-white" : "bg-gray-200"
                  }`}
                  onClick={() => setPage(pageNumber)}
                >
                  {pageNumber}
                </button>
              </li>
            );
          })}
        </ul>
        <button onClick={nextPage} className="text-sm">
          Next
        </button>
      </div>
      <div className="w-screen flex flex-col fixed bottom-6 items-center">
        <div
          className="bg-gray-800 rounded-lg flex flex-col items-center justify-center text-xl text-white font-medium font-serif w-60 h-14"
          onClick={() => {
            navigate("/add-citizen");
          }}
        >
          Add Citizen
        </div>
      </div>
    </section>
  );
}
