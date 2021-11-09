/*
SPDX-FileCopyrightText: 2021 SAP SE or an SAP affiliate company and Gardener contributors

SPDX-License-Identifier: Apache-2.0
*/

package target_test

import (
	"github.com/gardener/gardenctl-v2/pkg/target"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/pflag"
)

var _ = Describe("Target Flags", func() {
	It("should return an empty set of target flags", func() {
		tf := target.NewTargetFlags("", "", "", "")
		Expect(tf).NotTo(BeNil())
		Expect(tf.GardenName()).To(BeEmpty())
		Expect(tf.ProjectName()).To(BeEmpty())
		Expect(tf.SeedName()).To(BeEmpty())
		Expect(tf.ShootName()).To(BeEmpty())
		Expect(tf.IsTargetValid()).To(BeFalse())
	})

	It("should return valid set of target flags", func() {
		tf := target.NewTargetFlags("garden", "project", "", "shoot")
		Expect(tf).NotTo(BeNil())
		Expect(tf.GardenName()).To(Equal("garden"))
		Expect(tf.ProjectName()).To(Equal("project"))
		Expect(tf.SeedName()).To(BeEmpty())
		Expect(tf.ShootName()).To(Equal("shoot"))
		Expect(tf.IsTargetValid()).To(BeTrue())
	})

	It("should add target flags to a cobra FlagSet", func() {
		flags := &pflag.FlagSet{}
		tf := target.NewTargetFlags("", "", "", "")
		Expect(flags.HasFlags()).To(BeFalse())
		tf.AddFlags(flags)
		var names []string
		flags.VisitAll(func(flag *pflag.Flag) {
			names = append(names, flag.Name)
		})
		Expect(names).To(Equal([]string{"garden", "project", "seed", "shoot"}))
	})

	It("should validate target flags", func() {
		Expect(target.NewTargetFlags("", "project", "", "shoot").IsTargetValid()).To(BeFalse())
		Expect(target.NewTargetFlags("garden", "", "", "shoot").IsTargetValid()).To(BeTrue())
		Expect(target.NewTargetFlags("garden", "project", "seed", "shoot").IsTargetValid()).To(BeFalse())
		Expect(target.NewTargetFlags("garden", "", "", "").IsTargetValid()).To(BeTrue())
		Expect(target.NewTargetFlags("garden", "project", "", "").IsTargetValid()).To(BeTrue())
		Expect(target.NewTargetFlags("garden", "", "seed", "").IsTargetValid()).To(BeTrue())
		Expect(target.NewTargetFlags("garden", "project", "", "shoot").IsTargetValid()).To(BeTrue())
		Expect(target.NewTargetFlags("garden", "", "seed", "shoot").IsTargetValid()).To(BeTrue())
	})

	It("should override a target with target flags", func() {
		tf := target.NewTargetFlags("garden", "project", "", "shoot")
		t, err := tf.OverrideTarget(target.NewTarget("a", "b", "c", "d"))
		Expect(err).NotTo(HaveOccurred())
		Expect(t.GardenName()).To(Equal("garden"))
		Expect(t.ProjectName()).To(Equal("project"))
		Expect(t.SeedName()).To(BeEmpty())
		Expect(t.ShootName()).To(Equal("shoot"))
	})

	It("should fail to override a target", func() {
		tf := target.NewTargetFlags("", "", "", "shoot")
		_, err := tf.OverrideTarget(target.NewTarget("", "b", "c", "d"))
		Expect(err).To(HaveOccurred())
	})
})