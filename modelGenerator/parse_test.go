package modelGenerator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var edmxXmlString string = `This XML file does not appear to have any style information associated with it. The document tree is shown below.
<edmx:Edmx xmlns:edmx="http://docs.oasis-open.org/odata/ns/edmx" Version="4.0">
<edmx:DataServices>
<Schema xmlns="http://docs.oasis-open.org/odata/ns/edm" Namespace="Trippin">
<EntityType Name="Person">
<Key>
<PropertyRef Name="UserName"/>
</Key>
<Property Name="UserName" Type="Edm.String" Nullable="false"/>
<Property Name="FirstName" Type="Edm.String" Nullable="false"/>
<Property Name="LastName" Type="Edm.String" MaxLength="26"/>
<Property Name="MiddleName" Type="Edm.String"/>
<Property Name="Gender" Type="Trippin.PersonGender" Nullable="false"/>
<Property Name="Age" Type="Edm.Int64"/>
<Property Name="Emails" Type="Collection(Edm.String)"/>
<Property Name="AddressInfo" Type="Collection(Trippin.Location)"/>
<Property Name="HomeAddress" Type="Trippin.Location"/>
<Property Name="FavoriteFeature" Type="Trippin.Feature" Nullable="false"/>
<Property Name="Features" Type="Collection(Trippin.Feature)" Nullable="false"/>
<NavigationProperty Name="Friends" Type="Collection(Trippin.Person)"/>
<NavigationProperty Name="BestFriend" Type="Trippin.Person"/>
<NavigationProperty Name="Trips" Type="Collection(Trippin.Trip)"/>
</EntityType>
<EntityType Name="Airline">
<Key>
<PropertyRef Name="AirlineCode"/>
</Key>
<Property Name="AirlineCode" Type="Edm.String" Nullable="false"/>
<Property Name="Name" Type="Edm.String"/>
</EntityType>
<EntityType Name="Airport">
<Key>
<PropertyRef Name="IcaoCode"/>
</Key>
<Property Name="Name" Type="Edm.String"/>
<Property Name="IcaoCode" Type="Edm.String" Nullable="false"/>
<Property Name="IataCode" Type="Edm.String"/>
<Property Name="Location" Type="Trippin.AirportLocation"/>
</EntityType>
<ComplexType Name="Location">
<Property Name="Address" Type="Edm.String"/>
<Property Name="City" Type="Trippin.City"/>
</ComplexType>
<ComplexType Name="City">
<Property Name="Name" Type="Edm.String"/>
<Property Name="CountryRegion" Type="Edm.String"/>
<Property Name="Region" Type="Edm.String"/>
</ComplexType>
<ComplexType Name="AirportLocation" BaseType="Trippin.Location">
<Property Name="Loc" Type="Edm.GeographyPoint"/>
</ComplexType>
<ComplexType Name="EventLocation" BaseType="Trippin.Location">
<Property Name="BuildingInfo" Type="Edm.String"/>
</ComplexType>
<EntityType Name="Trip">
<Key>
<PropertyRef Name="TripId"/>
</Key>
<Property Name="TripId" Type="Edm.Int32" Nullable="false"/>
<Property Name="ShareId" Type="Edm.Guid" Nullable="false"/>
<Property Name="Name" Type="Edm.String"/>
<Property Name="Budget" Type="Edm.Single" Nullable="false"/>
<Property Name="Description" Type="Edm.String"/>
<Property Name="Tags" Type="Collection(Edm.String)"/>
<Property Name="StartsAt" Type="Edm.DateTimeOffset" Nullable="false"/>
<Property Name="EndsAt" Type="Edm.DateTimeOffset" Nullable="false"/>
<NavigationProperty Name="PlanItems" Type="Collection(Trippin.PlanItem)"/>
</EntityType>
<EntityType Name="PlanItem">
<Key>
<PropertyRef Name="PlanItemId"/>
</Key>
<Property Name="PlanItemId" Type="Edm.Int32" Nullable="false"/>
<Property Name="ConfirmationCode" Type="Edm.String"/>
<Property Name="StartsAt" Type="Edm.DateTimeOffset" Nullable="false"/>
<Property Name="EndsAt" Type="Edm.DateTimeOffset" Nullable="false"/>
<Property Name="Duration" Type="Edm.Duration" Nullable="false"/>
</EntityType>
<EntityType Name="Event" BaseType="Trippin.PlanItem">
<Property Name="OccursAt" Type="Trippin.EventLocation"/>
<Property Name="Description" Type="Edm.String"/>
</EntityType>
<EntityType Name="PublicTransportation" BaseType="Trippin.PlanItem">
<Property Name="SeatNumber" Type="Edm.String"/>
</EntityType>
<EntityType Name="Flight" BaseType="Trippin.PublicTransportation">
<Property Name="FlightNumber" Type="Edm.String"/>
<NavigationProperty Name="Airline" Type="Trippin.Airline"/>
<NavigationProperty Name="From" Type="Trippin.Airport"/>
<NavigationProperty Name="To" Type="Trippin.Airport"/>
</EntityType>
<EntityType Name="Employee" BaseType="Trippin.Person">
<Property Name="Cost" Type="Edm.Int64" Nullable="false"/>
<NavigationProperty Name="Peers" Type="Collection(Trippin.Person)"/>
</EntityType>
<EntityType Name="Manager" BaseType="Trippin.Person">
<Property Name="Budget" Type="Edm.Int64" Nullable="false"/>
<Property Name="BossOffice" Type="Trippin.Location"/>
<NavigationProperty Name="DirectReports" Type="Collection(Trippin.Person)"/>
</EntityType>
<EnumType Name="PersonGender">
<Member Name="Male" Value="0"/>
<Member Name="Female" Value="1"/>
<Member Name="Unknown" Value="2"/>
</EnumType>
<EnumType Name="Feature">
<Member Name="Feature1" Value="0"/>
<Member Name="Feature2" Value="1"/>
<Member Name="Feature3" Value="2"/>
<Member Name="Feature4" Value="3"/>
</EnumType>
<Function Name="GetPersonWithMostFriends">
<ReturnType Type="Trippin.Person"/>
</Function>
<Function Name="GetNearestAirport">
<Parameter Name="lat" Type="Edm.Double" Nullable="false"/>
<Parameter Name="lon" Type="Edm.Double" Nullable="false"/>
<ReturnType Type="Trippin.Airport"/>
</Function>
<Function Name="GetFavoriteAirline" IsBound="true" EntitySetPath="person">
<Parameter Name="person" Type="Trippin.Person"/>
<ReturnType Type="Trippin.Airline"/>
</Function>
<Function Name="GetFriendsTrips" IsBound="true">
<Parameter Name="person" Type="Trippin.Person"/>
<Parameter Name="userName" Type="Edm.String" Nullable="false" Unicode="false"/>
<ReturnType Type="Collection(Trippin.Trip)"/>
</Function>
<Function Name="GetInvolvedPeople" IsBound="true">
<Parameter Name="trip" Type="Trippin.Trip"/>
<ReturnType Type="Collection(Trippin.Person)"/>
</Function>
<Action Name="ResetDataSource"/>
<Action Name="UpdateLastName" IsBound="true">
<Parameter Name="person" Type="Trippin.Person"/>
<Parameter Name="lastName" Type="Edm.String" Nullable="false" Unicode="false"/>
<ReturnType Type="Edm.Boolean" Nullable="false"/>
</Action>
<Action Name="ShareTrip" IsBound="true">
<Parameter Name="personInstance" Type="Trippin.Person"/>
<Parameter Name="userName" Type="Edm.String" Nullable="false" Unicode="false"/>
<Parameter Name="tripId" Type="Edm.Int32" Nullable="false"/>
</Action>
<EntityContainer Name="Container">
<EntitySet Name="People" EntityType="Trippin.Person">
<NavigationPropertyBinding Path="BestFriend" Target="People"/>
<NavigationPropertyBinding Path="Friends" Target="People"/>
<NavigationPropertyBinding Path="Trippin.Employee/Peers" Target="People"/>
<NavigationPropertyBinding Path="Trippin.Manager/DirectReports" Target="People"/>
</EntitySet>
<EntitySet Name="Airlines" EntityType="Trippin.Airline">
<Annotation Term="Org.OData.Core.V1.OptimisticConcurrency">
<Collection>
<PropertyPath>Name</PropertyPath>
</Collection>
</Annotation>
</EntitySet>
<EntitySet Name="Airports" EntityType="Trippin.Airport"/>
<Singleton Name="Me" Type="Trippin.Person">
<NavigationPropertyBinding Path="BestFriend" Target="People"/>
<NavigationPropertyBinding Path="Friends" Target="People"/>
<NavigationPropertyBinding Path="Trippin.Employee/Peers" Target="People"/>
<NavigationPropertyBinding Path="Trippin.Manager/DirectReports" Target="People"/>
</Singleton>
<FunctionImport Name="GetPersonWithMostFriends" Function="Trippin.GetPersonWithMostFriends" EntitySet="People"/>
<FunctionImport Name="GetNearestAirport" Function="Trippin.GetNearestAirport" EntitySet="Airports"/>
<ActionImport Name="ResetDataSource" Action="Trippin.ResetDataSource"/>
</EntityContainer>
</Schema>
</edmx:DataServices>
</edmx:Edmx>`

var parsedEdmx edmxSchema
var edmxParseError error
var hasParsedEdmx bool = false

func getParsedEdmx() (edmxSchema, error) {
	if !hasParsedEdmx {
		parsedEdmx, edmxParseError = parseEdmx([]byte(edmxXmlString))
		hasParsedEdmx = true
	}
	return parsedEdmx, edmxParseError
}

func Test_Parse_edmx(t *testing.T) {
	edmx, err := getParsedEdmx()
	assert.NoError(t, err)
	assert.Equal(t, "Trippin", edmx.Namespace)

	personEntityType, ok := edmx.EntityTypes["Person"]
	assert.True(t, ok)
	assert.Equal(t, "Person", personEntityType.Name)
	assert.Len(t, personEntityType.Properties, 11)
	usernameProperty, ok := personEntityType.Properties["UserName"]
	assert.True(t, ok)
	assert.Equal(t, "Edm.String", usernameProperty.Type)
	assert.Equal(t, "string", usernameProperty.goType())
	lastNameProperty, ok := personEntityType.Properties["LastName"]
	assert.True(t, ok)
	assert.Equal(t, "Edm.String", lastNameProperty.Type)
	assert.Equal(t, "nullable.Nullable[string]", lastNameProperty.goType())

	peopleEntitySet := edmx.EntitySets["People"]
	assert.Equal(t, "Trippin.Person", peopleEntitySet.EntityType)
	assert.Equal(t, personEntityType, peopleEntitySet.getEntityType())
}
