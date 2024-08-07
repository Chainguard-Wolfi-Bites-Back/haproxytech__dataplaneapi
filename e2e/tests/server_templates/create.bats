#!/usr/bin/env bats
#
# Copyright 2021 HAProxy Technologies
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

load '../../libs/dataplaneapi'
load "../../libs/get_json_path"
load '../../libs/haproxy_config_setup'
load '../../libs/resource_client'
load '../../libs/version'

load 'utils/_helpers'

@test "server_templates: Add a new server template 1" {
  PARENT_NAME="test_backend"
  resource_post "$_BACKEND_BASE_PATH/$PARENT_NAME/server_templates" "data/first/post.json" "force_reload=true"
	assert_equal "$SC" 201

  resource_get "$_BACKEND_BASE_PATH/$PARENT_NAME/server_templates/first"
  assert_equal "$SC" 200

  assert_equal "first" "$(get_json_path "$BODY" '.prefix')"
  assert_equal "1-50" "$(get_json_path "$BODY" '.num_or_range')"
  assert_equal "first.com" "$(get_json_path "$BODY" '.fqdn')"
  assert_equal "443" "$(get_json_path "$BODY" '.port')"
}

@test "server_templates: Add a new server template 2" {
  PARENT_NAME="test_backend"
  resource_post "$_BACKEND_BASE_PATH/$PARENT_NAME/server_templates" "data/second/post.json" "force_reload=true"
	assert_equal "$SC" 201

  resource_get "$_BACKEND_BASE_PATH/$PARENT_NAME/server_templates/second"
  assert_equal "$SC" 200

  assert_equal "second" "$(get_json_path "$BODY" '.prefix')"
  assert_equal "1-3" "$(get_json_path "$BODY" '.num_or_range')"
  assert_equal "second.com" "$(get_json_path "$BODY" '.fqdn')"
  assert_equal "8080" "$(get_json_path "$BODY" '.port')"
}

@test "server_templates: Add a new server template when missing frontend" {
  PARENT_NAME="ghost"
  resource_post "$_BACKEND_BASE_PATH/$PARENT_NAME/server_templates" "data/second/post.json" "force_reload=true"
	assert_equal "$SC" 400
}
