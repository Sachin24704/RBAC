const OperationModel = ({ operation, onRemove }) => {
  return (
    <div className="bg-gray-200 text-gray-800 px-3 py-1 m-1 rounded flex items-center">
      <span>{operation.name}</span>
      <button onClick={onRemove} className="text-red-500 ml-2 font-bold">
        X
      </button>
    </div>
  );
};

export default OperationModel;
