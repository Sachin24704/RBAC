import ModelCard from "./ModelCard";

const RoleCard = ({role}) => {
  return (
    <div className="border p-4 mb-4 rounded">
      <h3 className="text-lg font-semibold mb-2">
        {role.name}
      </h3>
      {role.models.map((model) => (
        <ModelCard key ={model.id} roleId={role.id} model = {model} />
      ))}
    </div>
  )
}

export default RoleCard;

