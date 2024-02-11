import {CorsTests} from "../Home/types";

type AddTestPanelProps = {
    submission: () => void;
    onChange : (field:string, value:string) => void;
    state : CorsTests;
    cancel : () => void;
}

const AddTestPanel = (props:AddTestPanelProps) => {

    return (
        <div>
            <h1>Add Test Panel</h1>
            <button onClick={()=> props.cancel()}>Cancel</button>
            <input type="text" value={props.state.Origin} onChange={(e)=> props.onChange('Origin', e.target.value)} placeholder="Origin"/>
            <input type="text" onChange={(e)=> props.onChange('Endpoint', e.target.value)} placeholder="Endpoint"/>
            <input type="text" onChange={(e)=> props.onChange('Methods', e.target.value)} placeholder="Methods"/>
            <input type="text" onChange={(e)=> props.onChange('Headers', e.target.value)} placeholder="Headers"/>
            <input type="text" onChange={(e)=> props.onChange('Authentication', e.target.value)} placeholder="Authentication"/>
            <button onClick={props.submission}> SUBMIT TEST </button>
        </div>
    )
}

export default AddTestPanel