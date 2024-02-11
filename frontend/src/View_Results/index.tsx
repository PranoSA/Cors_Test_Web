import React, { useEffect } from 'react';

import {CorsTestResults} from '../Home/types'
import { useParams } from 'react-router-dom';

type FilterResultsQuery = {
    Origin : string,
    Desination : string,
    Method : string,
    Header : string,
}

const ViewResults: React.FC = () => {

    const [expanded, setExpanded] = React.useState<boolean[]>([]);

    const [results, setResults] = React.useState<CorsTestResults[]>([]);

    const [filterState, setFilterState] = React.useState<FilterResultsQuery>({
        Origin: '',
        Desination: '',
        Method: '',
        Header: '',
    })

    const [filteredResults, setFilteredResults] = React.useState<CorsTestResults[]>([])

    const {id} = useParams<{id:string}>();

    useEffect(() => {

        if (results.length === 0) {
            return
        }




            const corsTestResults = results.filter((result) => {
                if(filterState.Origin !== '' && result.Origin !== filterState.Origin) {
                    return false
                }
                return true 
            }).filter((result) => {
                if(filterState.Desination !== '' && result.Endpoint !== filterState.Desination) {
                    return false
                }
                return true 
            }
            ).filter((result) => {
                if(filterState.Method !== '' && result.Method !== filterState.Method) {
                    return false
                }
                return true 
            }
            ).filter((result) => {
                if(filterState.Header !== '' && result.Header !== filterState.Header) {
                    return false
                }
                return true
            })   
    
            console.log(corsTestResults)
    
            setFilteredResults(corsTestResults)



    }, [results, filterState])


    React.useEffect(() => {
        const fetchResults = async () => {
            try {
                const res = await fetch(`${import.meta.env.VITE_REACT_APP_API_URL}/results/${id}`,{
                    credentials: 'include'
                })
                const data = await res.json()
                setResults(data)
            }
            catch (error) {
                console.log(error)
            }
        }
        fetchResults()
    }, [])

    if (results.length === 0) {
        return (
            <div>
                <h1>No Results Found / Loading </h1>
            </div>
        )
    }

    //Filter Panel
    const FilterPanel = () => {
        return (
            <div>
                <input type="text" placeholder="Origin" value={filterState.Origin} onChange={(e) => setFilterState({...filterState, Origin: e.target.value})} />
                <input type="text" placeholder="Desination" value={filterState.Desination} onChange={(e) => setFilterState({...filterState, Desination: e.target.value})} />
                <input type="text" placeholder="Method" value={filterState.Method} onChange={(e) => setFilterState({...filterState, Method: e.target.value})} />
                <input type="text" placeholder="Header" value={filterState.Header} onChange={(e) => setFilterState({...filterState, Header: e.target.value})} />
            </div>
        )
    }

    return (
        <div>

            {FilterPanel()}
            {
                filteredResults.map((result, index) => {
                    return (
                        <div key={result.Id} className={`p-4 ${result.Okay ? 'bg-green-200' : 'bg-red-200'}`}>
                            <button className="br-400 p-5" onClick={()=> setExpanded([...expanded.slice(0,index), !expanded[index], ...expanded.slice(index+1,-1)])}> Expand </button>
                            <button className="bg-400 p-5" onClick={()=> setExpanded([...expanded.slice(0,index), !expanded[index], ...expanded.slice(index+1,-1)])}> Unexpand </button>
                            <h1 onClick={()=> setExpanded([...expanded, !expanded[index]])}>{result.Origin}:{result.Header}:{result.Method}</h1>
                            {
                                expanded[index] && (
                                    <div>
                                        <p>Endpoint: {result.Endpoint}</p>
                                        <p>Methods: {result.Method}</p>
                                        <p>Headers: {result.Header}</p>
                                        <p>Authentication: {result.Authentication}</p>
                                        <p>Okay: {result.Okay}</p>
                                        <p>Errors: {result.Errors}</p>
                                        <p>Access-Control-Allow-Origin: {result.Return_Access_Control_Allow_Origin}</p>
                                        <p>Access-Control-Allow-Methods: {result.Return_Access_Control_Allow_Method}</p>

                                        <p> </p>
                                    </div>
                                )
                            }
                        </div>
                    )
                })
            }
        </div>
    )
}

export default ViewResults;