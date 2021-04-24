// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Represents a 3D point in space.
// x is towards the right of the robot.
// y is towards the front of the robot.
// z is towards the top of the robot.
// Distances are in millimeters.
struct Point3D(f32, f32, f32);

pub enum Position {
    FrontRight,
    FrontLeft,
    BackRight,
    BackLeft,
}

enum Joint {
    BodyCoxa,
    CoxaFemur,
    FemurTibia,
}

const RAD_TO_DEG: f32 = 180.0 / 3.14159;

const COXA_LEN: f32 = 23.5;
const FEMUR_LEN: f32 = 38.0;
const TIBIA_LEN: f32 = 81.0;

// The leg represents a three-segment leg (coxa-femur-tibia), with the coxa
// joined to the body at the hip point and the tibia touching the floor at the
// toe point. The canonical toe point is (0,0,0), representing a leg that is
// pointing out from the body corner at a 45' angle, the femur horizontal and
// the tibia vertical.
pub struct Leg {
    hip_pt: Point3D,
    toe_pt: Point3D,
}

impl Leg {
    pub fn new(leg_pos: Position) -> Leg {
        // The canonical zero position of the toe is with the coxa at "45 degrees",
        // the femur horizontal, and the tibia vertical. Therefore the hip joint is
        // displaced by (coxa+femur)/sqrt(2), using Pythagoras' theorem.
        let hip_offset = (COXA_LEN + FEMUR_LEN) / f32::sqrt(2f32);
        Leg {
            hip_pt: match leg_pos {
                Position::FrontRight => Point3D(-hip_offset, -hip_offset, TIBIA_LEN),
                Position::FrontLeft => Point3D(hip_offset, -hip_offset, TIBIA_LEN),
                Position::BackRight => Point3D(-hip_offset, hip_offset, TIBIA_LEN),
                Position::BackLeft => Point3D(hip_offset, hip_offset, TIBIA_LEN),
            },
            toe_pt: Point3D(0f32, 0f32, 0f32),
        }
    }

    pub fn set_toe_point(&mut self, x: f32, y: f32, z: f32) {
        self.toe_pt = Point3D(x, y, z);
    }

    pub fn move_toe_point(&mut self, x: f32, y: f32, z: f32) {
        self.toe_pt = Point3D(x + self.toe_pt.0, y + self.toe_pt.1, z + self.toe_pt.2);
    }

    pub fn get_joint_angles(&self) -> (f32, f32, f32) {
        // Hip angle is measured counter-clockwise from a line projecting out from
        // the right side of the spider.
        let bc = f32::atan2(self.toe_pt.1 - self.hip_pt.1, self.toe_pt.0 - self.hip_pt.0);

        // Total horizontal distance from hip to toe.
        let total_horiz_reach = f32::sqrt(
            (self.toe_pt.0 - self.hip_pt.0) * (self.toe_pt.0 - self.hip_pt.0)
                + (self.toe_pt.1 - self.hip_pt.1) * (self.toe_pt.1 - self.hip_pt.1),
        );
        // Femur+tibia horizontal reach.
        let ft_horiz_reach = total_horiz_reach - COXA_LEN;
        // Femur+tibia reach in 3D space.
        // This gives us a triangle with sides (femur_len, tibia_len, ftReach).
        let ft_diag_reach = f32::sqrt(
            ft_horiz_reach * ft_horiz_reach
                + (self.toe_pt.2 - self.hip_pt.2) * (self.toe_pt.2 - self.hip_pt.2),
        );

        // Solve for angles, using the law of cosines.
        //   c^2 = a^2 + b^2 - 2*a*b*cos(C)
        //   2*a*b*cos(C) =  a^2 + b^2 - c^2
        //   cos(C) = (a^2 + b^2 - c^2) / (2*a*b)
        // Or in coding terms:
        //   cos_num = a*a + b*b - c*c
        //   cos_denom = 2*a*b
        //   angle_c = acosf(cos_num / cos_denom)

        // Coxa-Femur angle is measured counter-clockwise from horizontal, so up is
        // positive and down is negative. First, find the angle between the femur
        // and the imaginary line from the coxa-femur joint down to the  toe.
        let cos_num = ft_diag_reach * ft_diag_reach + FEMUR_LEN * FEMUR_LEN - TIBIA_LEN * TIBIA_LEN;
        let cos_denom = 2.0 * ft_diag_reach * FEMUR_LEN;
        let femur_reach_angle = f32::acos(cos_num / cos_denom);

        // Second, find the angle between horizontal and the imaginary line from the
        // coxa-femur joint down to the  toe.
        let horiz_reach_angle = f32::atan2(self.toe_pt.2 - self.hip_pt.2, ft_horiz_reach);
        let cf = femur_reach_angle + horiz_reach_angle;

        // Femur-Tibia angle is measured counter-clockwise from the femur, so it
        // will always be positive, and bigger numbers represent a further reach.
        let cos_num = FEMUR_LEN * FEMUR_LEN + TIBIA_LEN * TIBIA_LEN - ft_diag_reach * ft_diag_reach;
        let cos_denom = 2.0 * FEMUR_LEN * TIBIA_LEN;
        let ft = f32::acos(cos_num / cos_denom);

        // Convert radians to degrees.
        (bc * RAD_TO_DEG, cf * RAD_TO_DEG, ft * RAD_TO_DEG)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn is_near(a: f32, b: f32) -> bool {
        return (a - b).abs() < 0.25;
    }

    fn assert_tuple_approx_equal(a: (f32, f32, f32), b: (f32, f32, f32)) {
        assert!(is_near(a.0, b.0), "{} not near {}", a.0, b.0);
        assert!(is_near(a.1, b.1), "{} not near {}", a.1, b.1);
        assert!(is_near(a.2, b.2), "{} not near {}", a.2, b.2);
    }

    #[test]
    fn test_joint_angles_at_null_point() {
        let leg = Leg::new(Position::FrontRight);
        assert_tuple_approx_equal(leg.get_joint_angles(), (45.0, 0.0, 90.0));

        let leg = Leg::new(Position::FrontLeft);
        assert_tuple_approx_equal(leg.get_joint_angles(), (135.0, 0.0, 90.0));

        let leg = Leg::new(Position::BackRight);
        assert_tuple_approx_equal(leg.get_joint_angles(), (-45.0, 0.0, 90.0));

        let leg = Leg::new(Position::BackLeft);
        assert_tuple_approx_equal(leg.get_joint_angles(), (-135.0, 0.0, 90.0));
    }

    #[test]
    fn test_joint_angles_at_sides_of_body() {
        // Toe is to the side, pulled in a bit, and down.
        let hip_offset = (COXA_LEN + FEMUR_LEN) / f32::sqrt(2.0);
        let x = hip_offset / 2.0;
        let y = hip_offset;
        let z = -20.0;

        let mut leg = Leg::new(Position::FrontRight);
        leg.set_toe_point(-x, -y, z);
        assert_tuple_approx_equal(leg.get_joint_angles(), (0.0, -42.5, 111.0));

        let mut leg = Leg::new(Position::FrontLeft);
        leg.set_toe_point(x, -y, z);
        assert_tuple_approx_equal(leg.get_joint_angles(), (180.0, -42.5, 111.0));

        let mut leg = Leg::new(Position::BackRight);
        leg.set_toe_point(-x, y, z);
        assert_tuple_approx_equal(leg.get_joint_angles(), (0.0, -42.5, 111.0));

        let mut leg = Leg::new(Position::BackLeft);
        leg.set_toe_point(x, y, z);
        assert_tuple_approx_equal(leg.get_joint_angles(), (180.0, -42.5, 111.0));
    }

    #[test]
    fn test_joint_angles_ahead_and_behind_body() {
        // Toe to the front (or back), stretched out a bit, and above the hip.
        let hip_offset = (COXA_LEN + FEMUR_LEN) / f32::sqrt(2.0);
        let x = hip_offset;
        let y = -2.0/3.0 * hip_offset;
        let z = TIBIA_LEN + 10.0;

        let mut leg = Leg::new(Position::FrontRight);
        leg.set_toe_point(-x, -y, z);
        assert_tuple_approx_equal(leg.get_joint_angles(), (90.0, 145.0, 26.5));

        let mut leg = Leg::new(Position::FrontLeft);
        leg.set_toe_point(x, -y, z);
        assert_tuple_approx_equal(leg.get_joint_angles(), (90.0, 145.0, 26.5));

        let mut leg = Leg::new(Position::BackRight);
        leg.set_toe_point(-x, y, z);
        assert_tuple_approx_equal(leg.get_joint_angles(), (-90.0, 145.0, 26.5));

        let mut leg = Leg::new(Position::BackLeft);
        leg.set_toe_point(x, y, z);
        assert_tuple_approx_equal(leg.get_joint_angles(), (-90.0, 145.0, 26.5));
    }
}
