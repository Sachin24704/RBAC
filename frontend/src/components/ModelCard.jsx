"use client";
import { useState } from "react";
import OperationModel from "./OperationModel";

const ModelCard = ({ roleId, model }) => {
  const [operations, setOperations] = useState(model.operations);
  const availableOperations =
    model.name === "Leave"
      ? [
          { name: "Approve", id: 1 },
          { name: "Reject", id: 2 },
          { name: "Forward", id: 3 },
        ]
      : [
          { name: "View", id: 4 },
          { name: "Edit", id: 5 },
          { name: "Comment", id: 6 },
        ];

  const handleAddOperation = async (operationId) => {
    const opId = parseInt(operationId);
    const operation = availableOperations.find((op) => op.id === opId);
    if (!operation) return;
    try {
      await fetch(`http://localhost:8000/roles/${roleId}/operations`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          operation_ids: [operation.id],
        }),
      });
      setOperations([...operations, operation]);
    } catch (error) {
      console.error("error adding operations", error);
    }
  };

  const handleRemoveOperation = async (operationId) => {
    try {
      await fetch(`http://localhost:8000/roles/${roleId}/operations`, {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          operation_ids: [operationId],
        }),
      });
      setOperations(operations.filter((op) => op.id !== operationId));
    } catch (error) {
      console.error("error removing operations", error);
    }
  };
  return (
    <div className="border p-4 mb-4 rounded">
      <h4 className="text-md font-semibold mb-2">{model.name}</h4>
      <div className="mb-2">
        <select
          onChange={(e) => {
            if (e.target.value) {
              handleAddOperation(e.target.value);
              e.target.value = " ";
            }
          }}
          className="border p-2 rounded"
        >
          <option value=" ">Add Operations</option>
          {availableOperations
            .filter(
              (op) => !operations.some((operation) => operation.id === op.id)
            )
            .map((op) => (
              <option key={op.id} value={op.id}>
                {op.name}
              </option>
            ))}
        </select>
      </div>
      <div className="flex flex-wrap">
        {operations.map((operation) => (
          <OperationModel
            key={operation.id}
            operation={operation}
            onRemove={() => handleRemoveOperation(operation.id)}
          />
        ))}
      </div>
    </div>
  );
};

export default ModelCard;
