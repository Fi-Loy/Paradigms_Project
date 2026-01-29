// Project CSI2120/CSI2520
// Winter 2026
// Robert Laganiere, uottawa.ca

import java.util.ArrayList;

// this is the (incomplete) Program class
public class Program {
	
	private String programID;
	private String name;
	private int quota;
	private int[] rol;

	private ArrayList<Resident> matchedResidents = new ArrayList<>();


	// constructs a Program
    public Program(String id, String n, int q) {
	
		programID= id;
		name= n;
		quota= q;

	}

    // the rol in order of preference
	public void setROL(int[] rol) {
		
		this.rol= rol;
	}
	
	// string representation
	public String toString() {
      
       return "["+programID+"]: "+name+" {"+ quota+ "}" +" ("+rol.length+")";	  
	}

	/**
	 * Checks if resident is in ROL of current program.
	 * @param residentId
	 */
	public boolean member(int residentId){
		for (int id : rol) {
			if (residentId == id) {
				return true;
			}
		}
		return false;
	}

	/**
	 * Returns the rank of the resident within the 
	 * program ROL. Returns -1 if resident is not
	 * included in the ROL. 
	 */
	public int rank(int residentID){
		for (int i = 0; i < rol.length;i++) {
			if (rol[i] == residentID) {
				return i;
			}
		}
		return -1;
	}

	/**
	 * Returns the resident with the highest rank
	 * (lowest preference)
	 */
	public Resident leastPreferred() {
		Resident highestResident = null;
		int highestRank = -1;

		for (Resident resident : matchedResidents) {
			int currentRank = rank(resident.getResidentID());

			if (currentRank > highestRank) {
				highestRank = currentRank;
				highestResident = resident;
			}
		}
		return highestResident;
	}
	/**
	 * Adds Resident to the match list of the program if
	 * the program has not reached its quota or if the
	 * resident is preferred over an already matched
	 * resident. 
	 */
	public void addResident(Resident resident){
		if(quota == 0) return;
		int rank = rank(resident.getResidentID());
		if (rank == -1) return;


		if (matchedResidents.size() < quota) {
			matchedResidents.add(resident);
			resident.setMatchedProgram(this, rank);
			return;
		}
		Resident leastPreferred = leastPreferred();
		int leastPreferredRank = rank(leastPreferred.getResidentID());

		if (rank < leastPreferredRank) {
			matchedResidents.remove(leastPreferred);
			leastPreferred.clearMatch();
			matchedResidents.add(resident);
			resident.setMatchedProgram(this,rank);
		}
	}

}