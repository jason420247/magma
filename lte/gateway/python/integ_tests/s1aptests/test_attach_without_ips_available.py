"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

import unittest

import s1ap_types
import s1ap_wrapper


class TestAttachWithoutIpsAvailable(unittest.TestCase):

    def setUp(self):
        self._s1ap_wrapper = s1ap_wrapper.TestWrapper()
        self._blocks = []

    def tearDown(self):
        # TODO(wlq): Add back original blocks/restore state
        for block in self._blocks:
            self._s1ap_wrapper.mobility_util.add_ip_block(block)
        self._s1ap_wrapper.cleanup()

    def test_attach_without_ips_available(self):
        """ Attaching without available IPs in mobilityd """
        self._s1ap_wrapper.configUEDevice(1)

        # Clear blocks
        self._blocks.extend(self._s1ap_wrapper.mobility_util.list_ip_blocks())
        self._s1ap_wrapper.mobility_util.remove_ip_blocks(self._blocks)

        req = self._s1ap_wrapper.ue_req
        print("************************* Running End to End attach for ",
              "UE id ", req.ue_id)
        # Now actually attempt the attach
        self._s1ap_wrapper._s1_util.attach(
            req.ue_id, s1ap_types.tfwCmd.UE_END_TO_END_ATTACH_REQUEST,
            s1ap_types.tfwCmd.UE_ATTACH_REJECT_IND, s1ap_types.ueAttachFail_t)


if __name__ == "__main__":
    unittest.main()
