// Project CSI2120/CSI2520
// Winter 2026
// Robert Laganiere, uottawa.ca

// this is the (incomplete) Resident class
public class Resident {
	
	private int residentID;
	private String firstname;
	private String lastname;
	private String[] rol;
	private int rolIndex = 0;

	private Program matchedProgram;
	private int matchedRank;

	// constructs a Resident
    public Resident(int id, String fname, String lname) {
		residentID= id;
		firstname= fname;
		lastname= lname;
		matchedProgram = null;
		matchedRank = -1;
	}

	// getters
	public int getResidentID() {return residentID;}
	public String[] getRol() {return rol;}
	public String getLastname() {return lastname;}
	public Program getMatchedProgram() {return matchedProgram;}
	public int getMatchedRank() {return matchedRank;}

	// setters
	/** 
	 * Sets the Program of the Resident
	*/
	public void setMatchedProgram(Program p, int rank) {
		matchedProgram = p;
		matchedRank = rank;
	}
	/**checks if the resident is matched
	*/ 
	public boolean isMatched() {
		return matchedProgram != null;
	}

	/**
	 * Clears the match of the Resident
	 */
	public void clearMatch() {
		matchedProgram = null;
		matchedRank = -1;
	}

	/**
	 * Sets the rol in order of preference
	 */ 
	public void setROL(String[] rol) {
		this.rol= rol;
	}

	/**
	 * Checks if the resident has reached the end of his
	 * ROL and if there are available programs to match to.
	 */
	public boolean hasRol() {
		return (rolIndex < rol.length);
	}

	/**
	 * Returns the next listing in the resident ROL and
	 * increments the ROL index. Returns null if at the 
	 * end of the ROL.
	 */
	public String nextRol() {
		if (hasRol()) {
			String preferred = rol[rolIndex];
			rolIndex++;
			return preferred;
		}
		return null;
	}

	/**
	 * string representation
	 */
	public String toString() {
      
       return "["+residentID+"]: "+firstname+" "+ lastname+" ("+rol.length+")";
	}
	
	public String matchedString() {
		return lastname + "," + firstname + "," + residentID + "," +
				matchedProgram.getProgramID() + "," + matchedProgram.getName();

	}
	public String unmachtedString() {
		return lastname + "," + firstname + "," + residentID + "," +
				"XXX,NOT_MATCHED";
	}
}