# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: agent_api.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "agentapi.AgentApiRequest" do
    optional :node_key, :string, 1
    optional :agent_version, :string, 2
  end
  add_message "agentapi.AgentApiResponse" do
    optional :message, :string, 1
    optional :error_code, :string, 2
    optional :node_invalid, :bool, 3
  end
  add_message "agentapi.EnrollmentRequest" do
    optional :enroll_secret, :string, 1
    optional :host_identifier, :string, 2
  end
  add_message "agentapi.EnrollmentResponse" do
    optional :node_key, :string, 1
    optional :node_invalid, :bool, 2
  end
  add_message "agentapi.ConfigResponse" do
    optional :config_json_blob, :string, 1
    optional :node_invalid, :bool, 2
  end
  add_message "agentapi.LogCollection" do
    optional :node_key, :string, 1
    optional :log_type, :enum, 2, "agentapi.LogCollection.LogType"
    repeated :logs, :message, 3, "agentapi.LogCollection.Log"
  end
  add_message "agentapi.LogCollection.Log" do
    optional :data, :string, 1
  end
  add_enum "agentapi.LogCollection.LogType" do
    value :RESULT, 0
    value :STATUS, 1
    value :AGENT, 2
  end
  add_message "agentapi.QueryCollection" do
    repeated :queries, :message, 1, "agentapi.QueryCollection.Query"
    optional :node_invalid, :bool, 2
  end
  add_message "agentapi.QueryCollection.Query" do
    optional :id, :int32, 1
    optional :query, :string, 2
  end
  add_message "agentapi.ResultCollection" do
    optional :node_key, :string, 1
    repeated :results, :message, 2, "agentapi.ResultCollection.Result"
  end
  add_message "agentapi.ResultCollection.Result" do
    optional :id, :int32, 1
    repeated :rows, :message, 2, "agentapi.ResultCollection.Result.ResultRow"
    optional :status, :int32, 3
  end
  add_message "agentapi.ResultCollection.Result.ResultRow" do
    optional :id, :int32, 1
    repeated :columns, :message, 2, "agentapi.ResultCollection.Result.ResultRow.Column"
  end
  add_message "agentapi.ResultCollection.Result.ResultRow.Column" do
    optional :name, :string, 1
    optional :value, :string, 2
  end
end

module Agentapi
  AgentApiRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("agentapi.AgentApiRequest").msgclass
  AgentApiResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("agentapi.AgentApiResponse").msgclass
  EnrollmentRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("agentapi.EnrollmentRequest").msgclass
  EnrollmentResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("agentapi.EnrollmentResponse").msgclass
  ConfigResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("agentapi.ConfigResponse").msgclass
  LogCollection = Google::Protobuf::DescriptorPool.generated_pool.lookup("agentapi.LogCollection").msgclass
  LogCollection::Log = Google::Protobuf::DescriptorPool.generated_pool.lookup("agentapi.LogCollection.Log").msgclass
  LogCollection::LogType = Google::Protobuf::DescriptorPool.generated_pool.lookup("agentapi.LogCollection.LogType").enummodule
  QueryCollection = Google::Protobuf::DescriptorPool.generated_pool.lookup("agentapi.QueryCollection").msgclass
  QueryCollection::Query = Google::Protobuf::DescriptorPool.generated_pool.lookup("agentapi.QueryCollection.Query").msgclass
  ResultCollection = Google::Protobuf::DescriptorPool.generated_pool.lookup("agentapi.ResultCollection").msgclass
  ResultCollection::Result = Google::Protobuf::DescriptorPool.generated_pool.lookup("agentapi.ResultCollection.Result").msgclass
  ResultCollection::Result::ResultRow = Google::Protobuf::DescriptorPool.generated_pool.lookup("agentapi.ResultCollection.Result.ResultRow").msgclass
  ResultCollection::Result::ResultRow::Column = Google::Protobuf::DescriptorPool.generated_pool.lookup("agentapi.ResultCollection.Result.ResultRow.Column").msgclass
end