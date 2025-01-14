#!/usr/bin/env python3
# Copyright 2021 The Kubeflow Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import kfp as kfp
from .use_run_id import pipeline_use_run_id
from kfp.samples.test.utils import run_pipeline_func, TestCase

run_pipeline_func([
    TestCase(
        pipeline_func=pipeline_use_run_id,
        mode=kfp.dsl.PipelineExecutionMode.V2_ENGINE,
    ),
])