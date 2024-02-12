import React, { useEffect } from 'react';

import { CorsTestResults } from '../Home/types';
import { useParams } from 'react-router-dom';

type FilterResultsQuery = {
  Origin: string;
  Desination: string;
  Method: string;
  Header: string;
};

const ViewResults: React.FC = () => {
  const [expanded, setExpanded] = React.useState<boolean[]>([]);

  const [results, setResults] = React.useState<CorsTestResults[]>([]);

  const [filterState, setFilterState] = React.useState<FilterResultsQuery>({
    Origin: '',
    Desination: '',
    Method: '',
    Header: '',
  });

  const [filteredResults, setFilteredResults] = React.useState<
    CorsTestResults[]
  >([]);

  const { id } = useParams<{ id: string }>();

  useEffect(() => {
    if (results.length === 0) {
      return;
    }

    const corsTestResults = results
      .filter((result) => {
        if (filterState.Origin !== '' && result.Origin !== filterState.Origin) {
          return false;
        }
        return true;
      })
      .filter((result) => {
        if (
          filterState.Desination !== '' &&
          result.Endpoint !== filterState.Desination
        ) {
          return false;
        }
        return true;
      })
      .filter((result) => {
        if (filterState.Method !== '' && result.Method !== filterState.Method) {
          return false;
        }
        return true;
      })
      .filter((result) => {
        if (filterState.Header !== '' && result.Header !== filterState.Header) {
          return false;
        }
        return true;
      });

    console.log(corsTestResults);

    setFilteredResults(corsTestResults);
  }, [results, filterState]);

  React.useEffect(() => {
    const fetchResults = async () => {
      try {
        const res = await fetch(
          `${import.meta.env.VITE_REACT_APP_API_URL}/results/${id}`,
          {
            credentials: 'include',
          }
        );
        const data = await res.json();
        setResults(data);
      } catch (error) {
        console.log(error);
      }
    };
    fetchResults();
  }, []);

  if (results.length === 0) {
    return (
      <div>
        <h1>No Results Found / Loading </h1>
      </div>
    );
  }

  const getAllOrigins = () => {
    const origins = results.map((result) => result.Origin);

    const seen = new Set();

    const filteredOrigins = origins.filter((el) => {
      const duplicate = seen.has(el);
      seen.add(el);
      return !duplicate;
    });

    return filteredOrigins;
  };

  const getAllDestinations = () => {
    const origins = results.map((result) => result.Endpoint);

    const seen = new Set();

    const filteredOrigins = origins.filter((el) => {
      const duplicate = seen.has(el);
      seen.add(el);
      return !duplicate;
    });

    return filteredOrigins;
  };

  const getAllHeaders = () => {
    const origins = results.map((result) => result.Header);

    const seen = new Set();

    const filteredOrigins = origins.filter((el) => {
      const duplicate = seen.has(el);
      seen.add(el);
      return !duplicate;
    });

    return filteredOrigins;
  };

  const methods: string[] = [
    'GET',
    'POST',
    'DELETE',
    'PUT',
    'PATCH',
    'OPTIONS',
    'HEAD',
    'CONNECT',
    'TRACE',
    'COPY',
    'LOCK',
    'MKCOL',
    'MOVE',
    'PROPFIND',
    'PROPPATCH',
    'SEARCH',
    'UNLOCK',
    'BIND',
    'REBIND',
    'UNBIND',
    'ACL',
    'REPORT',
    'MKACTIVITY',
    'CHECKOUT',
    'MERGE',
    'M-SEARCH',
    'NOTIFY',
    'SUBSCRIBE',
    'UNSUBSCRIBE',
    'PATCH',
    'PURGE',
    'MKCALENDAR',
    'LINK',
    'UNLINK',
    'SOURCE',
    'RAW',
    'VIEW',
  ];

  //Filter Panel
  const FilterPanel = () => {
    return (
      <div className="justify-around flex flex-row w-full p-4 bg-gray-200">
        <div>
          <label htmlFor="origins">Origin</label>
          <input
            list="origins"
            type="text"
            placeholder="Origin"
            value={filterState.Origin}
            onChange={(e) =>
              setFilterState({ ...filterState, Origin: e.target.value })
            }
          />
          <datalist id="origins">
            {getAllOrigins().map((origin, index) => {
              return <option key={index} value={origin} />;
            })}
          </datalist>
        </div>

        <div>
          <label htmlFor="destination">Destination</label>
          <input
            list="destination"
            type="text"
            placeholder="Destination"
            value={filterState.Desination}
            onChange={(e) =>
              setFilterState({ ...filterState, Desination: e.target.value })
            }
          />
          <datalist id="destination">
            {getAllDestinations().map((origin, index) => {
              return <option key={index} value={origin} />;
            })}
          </datalist>
        </div>

        <div>
          <label htmlFor="method">Method</label>
          <input
            list="method"
            type="text"
            placeholder="Method"
            value={filterState.Method}
            onChange={(e) =>
              setFilterState({ ...filterState, Method: e.target.value })
            }
          />
          <datalist id="method">
            {methods.map((origin, index) => {
              return <option key={index} value={origin} />;
            })}
          </datalist>
        </div>

        <div>
          <label htmlFor="header">Header</label>
          <input
            list="header"
            type="text"
            placeholder="Header"
            value={filterState.Header}
            onChange={(e) =>
              setFilterState({ ...filterState, Header: e.target.value })
            }
          />
          <datalist id="header">
            {getAllHeaders().map((origin, index) => {
              return <option key={index} value={origin} />;
            })}
          </datalist>
        </div>
      </div>
    );
  };

  return (
    <div className="w-full flex flex-wrap">
      {FilterPanel()}
      {filteredResults.map((result, index) => {
        return (
          <div
            key={result.Id}
            className={`p-4 ${
              result.Okay ? 'bg-green-200' : 'bg-red-200'
            } border-5 w-full md:w-1/2`}
            w-full
          >
            <button
              className="br-400 p-5"
              onClick={() =>
                setExpanded([
                  ...expanded.slice(0, index),
                  !expanded[index],
                  ...expanded.slice(index + 1),
                ])
              }
            >
              {' '}
              Expand{' '}
            </button>
            <button
              className="bg-400 p-5"
              onClick={() =>
                setExpanded([
                  ...expanded.slice(0, index),
                  !expanded[index],
                  ...expanded.slice(index + 1),
                ])
              }
            >
              {' '}
              Unexpand{' '}
            </button>
            <div className="w-full flex flex-wrap border-5">
              <div className="w-1/2">Origin</div>
              <div className="w-1/2">{result.Origin}</div>
            </div>
            <div className="w-full flex flex-wrap border-5">
              <div className="w-1/2">Endpoint</div>
              <div className="w-1/2">{result.Endpoint}</div>
            </div>
            <div className="w-full flex flex-wrap border-5">
              <div className="w-1/2">Method</div>
              <div className="w-1/2">{result.Method}</div>
            </div>
            <div className="w-full flex flex-wrap border-5">
              <div className="w-1/2">Header</div>
              <div className="w-1/2">{result.Header}</div>
            </div>
            {expanded[index] && (
              <div>
                <p>Endpoint: {result.Endpoint}</p>
                <p>Methods: {result.Method}</p>
                <p>Headers: {result.Header}</p>
                <p>Authentication: {result.Authentication}</p>
                <p>Okay: {result.Okay}</p>
                <p>Errors: {result.Errors}</p>
                <p>
                  Access-Control-Allow-Origin:{' '}
                  {result.Return_Access_Control_Allow_Origin}
                </p>
                <p>
                  Access-Control-Allow-Methods:{' '}
                  {result.Return_Access_Control_Allow_Method}
                </p>

                <p> </p>
              </div>
            )}
          </div>
        );
      })}
    </div>
  );
};

export default ViewResults;
