import { CorsTests } from '../Home/types';

type AddTestPanelProps = {
  submission: () => void;
  onChange: (field: string, value: string) => void;
  state: CorsTests;
  cancel: () => void;
  edit: boolean;
};

const AddTestPanel = (props: AddTestPanelProps) => {
  return (
    <div className="w-full flex flex-wrap justify-center text-center p-40">
      <div className="w-full h-1 bg-gray-200 text-center">
        <h1 className="text-center">
          {' '}
          {props.edit ? 'Edit Test' : 'Add Test'}
        </h1>
      </div>
      <div className="w-full h-1 bg-gray-200 text-center">
        <button className="p-5 bg-red-250" onClick={() => props.cancel()}>
          Cancel
        </button>
      </div>

      <div className="w-full flex">
        <input
          type="text"
          value={props.state.Origin}
          onChange={(e) => props.onChange('Origin', e.target.value)}
          placeholder="Origin"
        />
      </div>
      <div className="w-full flex">
        <input
          type="text"
          onChange={(e) => props.onChange('Endpoint', e.target.value)}
          placeholder="Endpoint"
        />
      </div>
      <div className="w-full flex">
        <input
          type="text"
          onChange={(e) => props.onChange('Methods', e.target.value)}
          placeholder="Methods"
        />
      </div>
      <div className="w-full flex">
        <input
          type="text"
          onChange={(e) => props.onChange('Headers', e.target.value)}
          placeholder="Headers"
        />
      </div>
      <div className="w-full flex">
        <input
          type="text"
          onChange={(e) => props.onChange('Authentication', e.target.value)}
          placeholder="Authentication"
        />
      </div>
      <div className="w-full flex">
        <button onClick={props.submission}> SUBMIT TEST </button>
      </div>
      <div>
        <button onClick={() => props.cancel()}>Cancel</button>
      </div>
    </div>
  );
};

export default AddTestPanel;
