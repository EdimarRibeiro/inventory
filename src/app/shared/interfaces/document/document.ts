
import { Participant } from "@interfaces/general/participant";
import { Inventory } from "@interfaces/inventory/inventory";

export interface Document {
    id: number;
    InventoryId   :number
	Inventory     :Inventory
	ParticipantId :number
	Participant   :Participant
	OperationId   :string
	EmitentTypeId :string
	ModelId       :string
	StatusId      :string
	Serie         :string
	Number        :string
	DocumentKey   :string
	EmitentDate   :Date
	ExitDate      :Date
	DocumentValue :number
	PayTypeId     :string
	Discount      :number
	Reduction     :number
	ProductValue  :number
	FreightType   :string
	FreightValue  :number
	SafeValue     :number
	ExpenseValue  :number
	BaseIcms      :number
	ValueIcms     :number
	BaseIcmsSt    :number
	ValueIcmsSt   :number
	ValueIpi      :number
	ValuePis      :number
	ValueCofins   :number
	ValuePisSt    :number
	ValueCofinsSt :number
	Origined      :string
	Processed     :boolean
	Imported      :boolean
	Cloused       :boolean
}